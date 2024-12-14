package middleware

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"net/http"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
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
