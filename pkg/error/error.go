package errorhandler

import (
	dto "kredit-plus/internal/dto/result"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(c echo.Context, err error, message string, httpStatus int) error {
	// Log the error with context
	c.Logger().Errorf("Error: %v, Message: %s, HTTP Status: %d", err, message, httpStatus)

	// Return JSON response with structured error details
	return c.JSON(httpStatus, dto.ErrorResult{
		Status:  httpStatus,
		Message: message,
	})
}
