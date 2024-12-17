package middleware

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"net/http"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/labstack/echo/v4"
	"io"
	"os"
	"path/filepath"
	"time"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		file, err := c.FormFile("file")
		if err != nil {
			return fmt.Errorf("Image not found. Please upload a valid file.")
		}

		src, err := file.Open()
		if err != nil {
			return fmt.Errorf("Failed to open the uploaded file.")
		}
		defer src.Close()

		if file.Size > 5*1024*1024 {
			return fmt.Errorf("file size exceeds the 5MB limit")
		}
		if _, err := os.Stat("public"); os.IsNotExist(err) {
			err := os.Mkdir("public", os.ModePerm)
			if err != nil {
				return fmt.Errorf("Failed to create 'public' directory.") // Handle error
			}
		}
		buf := make([]byte, 512)
		_, err = src.Read(buf)
		if err != nil && err.Error() != "EOF" {
			return err
		}

		// Detect MIME type
		mimeType := http.DetectContentType(buf)

		if mimeType != "image/jpeg" && mimeType != "image/png" {
			return fmt.Errorf("invalid file type: %s. Only JPEG and PNG are allowed", mimeType)
		}
		// Create a destination file
		extension := filepath.Ext(file.Filename)
		fileNameWithoutExt := file.Filename[:len(file.Filename)-len(extension)]
		timestamp := fmt.Sprintf("%d", time.Now().Unix()) // Convert Unix timestamp to string
		newFileName := fmt.Sprintf("public/%s%s", fileNameWithoutExt+"-"+timestamp, extension)

		dst, err := os.Create(newFileName)
		if err != nil {
			return err
		}
		defer dst.Close()
		// Reset the file reader and copy to the destination
		_, err = src.Seek(0, io.SeekStart)
		if err != nil {
			return err
		}

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		return next(c)
	}
}
func UploadFileDisintegration(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		form, err := c.MultipartForm()
		if err != nil {
			return echo.ErrBadRequest
		}

		files, ok := form.File["image"]
		if !ok {
			c.Set("dataFile", make(map[string][]byte))
			return next(c)
		}
		// var myMap map[string][]byte
		myMap := make(map[string][]byte)

		Counter := 0

		for _, file := range files {
			// Extract file extension

			imageSize := resizeImage(file.Size)

			imageBytes, _ := file.Open()

			defer imageBytes.Close()

			buffer := make([]byte, file.Size)
			_, err = imageBytes.Read(buffer)

			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error reading uploaded file data: %v", err))
			}

			img, err := imaging.Decode(bytes.NewReader(buffer), imaging.AutoOrientation(true))

			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid image format: %v", err))
			}

			srcs := imaging.Resize(img, imageSize, 0, imaging.Lanczos)

			buf := new(bytes.Buffer)
			err = jpeg.Encode(buf, srcs, nil)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error encoding resized image: %v", err))
			}

			datafile := buf.Bytes()

			Counter = Counter + 1

			TotalName := strconv.Itoa(Counter)

			myMap["image"+TotalName] = datafile

		}
		c.Set("dataFile", myMap)
		return next(c)
	}
}

func resizeImage(fileSize int64) int {

	if fileSize == 0 {
		return 0
	}

	if fileSize > 10*1000000 {
		return 500
	}
	return 800
}
