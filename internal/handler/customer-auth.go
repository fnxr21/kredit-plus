package handler

import (
	authdto "kredit-plus/internal/dto/auth"
	customerdto "kredit-plus/internal/dto/customer.go"
	dto "kredit-plus/internal/dto/result"
	repositories "kredit-plus/internal/repository"
	"kredit-plus/pkg/bcrypt"
	errorhandler "kredit-plus/pkg/error"
	jwtToken "kredit-plus/pkg/jwt"
	"log"
	"net/http"

	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerCustomerAuth struct {
	CustomerAuthRepository repositories.CustomerAuth
}

func HandlerCustomerAuth(CustomerAuthRepository repositories.CustomerAuth) *handlerCustomerAuth {
	return &handlerCustomerAuth{CustomerAuthRepository}
}

func (h *handlerCustomerAuth) LoginCustomer(c echo.Context) error {
	// catch json  or form type
	request := new(authdto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	// Step 2: Bind the incoming JSON payload to the LoginRequest object.
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	error := c.Validate(request)

	if error != nil {
		return errorhandler.ErrorHandler(c, error, error.Error(), http.StatusBadRequest)
	}

	customer, err := h.CustomerAuthRepository.LoginCustomer(request.Username)
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
	}
	// Compare the provided password with the stored password hash using bcrypt.
	isValid := bcrypt.CheckPasswordHash(request.Password, customer.Password)

	if !isValid {
		return errorhandler.ErrorHandler(c, err, "Incorrect Password", http.StatusBadRequest)
	}

	//generate a JWT token with the user's claims.
	claims := jwt.MapClaims{}
	claims["id"] = customer.ID
	claims["name"] = customer.Username
	claims["status"] = "customer"
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // Set token expiration to 2 hours from now.

	//Generate the JWT token using the claims.
	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	response := authdto.LoginResponse{
		Token: token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: response})
}

func (h *handlerCustomerAuth) ReauthCustomer(c echo.Context) error {
	customerLogin := c.Get("customerLogin")
	customerID := customerLogin.(jwt.MapClaims)["id"].(float64)

	user, err := h.CustomerAuthRepository.ReauthCustomer(uint(customerID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "customer-not-found", http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, dto.SuccessReauth{Status: http.StatusOK, Data: user.Username + " " + "still-active"})

}

func (h *handlerAdminAuth) RegisterCustomer(c echo.Context) error {
	request := new(customerdto.RequestRegisterCustomer)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	// Step 2: Bind the incoming JSON payload to the LoginRequest object.
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	error := c.Validate(request)

	if error != nil {
		return errorhandler.ErrorHandler(c, error, error.Error(), http.StatusBadRequest)
	}
	pass, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return errorhandler.ErrorHandler(c, error, error.Error(), http.StatusBadRequest)
	}
	// admin := models.MyUser{
	// 	Username:    request.Username,
	// 	Password:    pass,
	// 	PhoneNumber: request.PhoneNumber,
	// 	Email:       request.Email,
	// }
	// user, err := h.AdminAuthRepository.Register(admin)
	// if err != nil {
	// 	return errorhandler.ErrorHandler(c, err, "Register Failed", http.StatusUnauthorized)
	// }

	return c.JSON(http.StatusOK, dto.SuccessReauth{Status: http.StatusOK, Data: user.Username + " " + "Register"})

}
