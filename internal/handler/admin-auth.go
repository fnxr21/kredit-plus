package handler

import (
	// admindto "kredit-plus/internal/dto/admin"
	"fmt"
	admindto "kredit-plus/internal/dto/admin"
	authdto "kredit-plus/internal/dto/auth"
	dto "kredit-plus/internal/dto/result"

	// "log"
	"time"

	// "kredit-plus/internal/models"
	// repositories "kredit-plus/internal/repository"
	service "kredit-plus/internal/services.go"
	// "kredit-plus/pkg/bcrypt"
	errorhandler "kredit-plus/pkg/error"
	// jwtToken "kredit-plus/pkg/jwt"
	// "log"
	"net/http"

	// "time"
	// "github.com/go-playground/validator/v10"

	// "github.com/golang-jwt/jwt/v4"
	validator "github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerAdminAuth struct {
	AdminAuthService service.AuthService
}

func HandlerAdminAuthtest(AdminAuthService service.AuthService) *handlerAdminAuth {
	return &handlerAdminAuth{AdminAuthService}
}

func (r *handlerAdminAuth) LoginAdmin(c echo.Context) error {

	request := new(authdto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	err := c.Validate(request)
	// validator.ValidationErrors
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			switch err.Tag() {
			case "required":
				return errorhandler.HandlerValidationError(c, http.StatusBadRequest, fmt.Sprintf("%s is required", err.Field()))
			case err.Tag():
				return errorhandler.HandlerValidationError(c, http.StatusBadRequest, fmt.Sprintf("%s is not valid "+err.Tag(), err.Field()))
			case "gte":
				errorhandler.HandlerValidationError(c, http.StatusBadRequest, fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param()))
			case "lte":
				errorhandler.HandlerValidationError(c, http.StatusBadRequest, fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param()))
			}
		}
	}

	token, err := r.AdminAuthService.LoginAdmin(request)

	if err != nil {
		return errorhandler.ErrorHandler(c, err, err.Error(), http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: token})
}

func (h *handlerAdminAuth) ReauthAdmin(c echo.Context) error {
	adminLogin := c.Get("adminLogin")
	adminID := adminLogin.(jwt.MapClaims)["id"].(float64)

	user, err := h.AdminAuthService.ReauthAdmin(uint(adminID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, dto.SuccessReauth{Status: http.StatusOK, Data: user + " " + "Still Active"})

}

func (h *handlerAdminAuth) RegisterAdmin(c echo.Context) error {
	request := new(admindto.RequestRegisterAdmin)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	error := c.Validate(request)
	if error != nil {
		return errorhandler.ErrorHandler(c, error, error.Error(), http.StatusBadRequest)
	}

	user, err := h.AdminAuthService.RegisterAdmin(request)
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, dto.SuccessReauth{Status: http.StatusOK, Data: user + " " + "Register"})

}

func (h *handlerAdminAuth) LogoutAdmin(c echo.Context) error {
	delete := &http.Cookie{
		Name:     "Auth",
		Value:    "none",
		Expires:  time.Now(),
		Path:     "/",
		HttpOnly: true,
	}
	c.SetCookie(delete)

	return c.JSON(http.StatusOK, dto.SuccessReauth{Status: http.StatusOK, Data: "Admin logged out successfully"})

}
