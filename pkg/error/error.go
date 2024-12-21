package errorhandler

import (
	"fmt"
	// "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	dto "kredit-plus/internal/dto/result"
)

// lama
func ErrorHandler(c echo.Context, err error, message string, httpStatus int) error {
	// Log the error with context
	c.Logger().Errorf("Error: %v, Message: %s, HTTP Status: %d", err, message, httpStatus)

	// Return JSON response with structured error details
	return c.JSON(httpStatus, dto.ErrorResult{
		Status:  httpStatus,
		Message: message,
	})
}

type ServiceError struct {
	Code    string
	Message string
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s", e.Code, e.Message)
}

func NewServiceError(code string, err error) *ServiceError {
	return &ServiceError{
		Code:    code,
		Message: err.Error(),
	}
}

// func ErrorFromValidationError(code string, err validator.ValidationErrors) *ServiceError {
// 	errMsg := err.Error()

//		return &ServiceError{
//			Code:    code,
//			Message: errMsg,
//		}
//	}
func HandlerValidationError(c echo.Context, code int, err string) error {

	return c.JSON(code, dto.ErrorResult{
		Status:  code,
		Message: err,
	})

}

type check struct {
}
