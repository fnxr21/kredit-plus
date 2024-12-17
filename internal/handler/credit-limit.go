package handler

import (
	creditlimitdto "kredit-plus/internal/dto/credit-limit"
	dto "kredit-plus/internal/dto/result"
	"kredit-plus/internal/models"
	repositories "kredit-plus/internal/repository"
	errorhandler "kredit-plus/pkg/error"
	// "log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerCreditLimit struct {
	CreditLimitRepository  repositories.CreditLimit
	AdminAuthRepository    repositories.AdminAuth
	CustomerAuthRepository repositories.CustomerAuth
}

func HandlerCreditLimit(
	CreditLimitRepository repositories.CreditLimit,
	AdminAuthRepository repositories.AdminAuth,
	CustomerAuthRepository repositories.CustomerAuth) *handlerCreditLimit {
	return &handlerCreditLimit{
		CreditLimitRepository:  CreditLimitRepository,
		AdminAuthRepository:    AdminAuthRepository,
		CustomerAuthRepository: CustomerAuthRepository}
}

func (h *handlerCreditLimit) CreateCreditLimit(c echo.Context) error {
	accessLogin := c.Get("adminLogin").(jwt.MapClaims)
	accessLoginID := accessLogin["id"].(float64)

	_, err := h.AdminAuthRepository.Reauth(uint(accessLoginID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
	}
	// catch json  type
	request := new(creditlimitdto.RequestRegisterCustomer)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	// Step 2: Bind the incoming JSON payload to the.
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	error := c.Validate(request)

	if error != nil {
		return errorhandler.ErrorHandler(c, error, error.Error(), http.StatusBadRequest)
	}

	tenorMonths, _ := strconv.ParseFloat(strings.TrimSpace(request.TenorMonths), 64)
	limitAmount, _ := strconv.ParseFloat(strings.TrimSpace(request.LimitAmount), 64)
	limit := models.CreditLimit{
		TenorMonths: tenorMonths,
		LimitAmount: limitAmount,
	}

	creditlimit, err := h.CreditLimitRepository.CreateCreditLimit(limit)
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
}

func (h *handlerCreditLimit) ListCreditLimit(c echo.Context) error {
	accessLogin, ok := c.Get("adminLogin").(jwt.MapClaims)
	// var accessLoginID float64
	if !ok || accessLogin == nil {
		userLogin, _ := c.Get("customerLogin").(jwt.MapClaims)
		if userLogin != nil {

			accessLoginID := userLogin["id"].(float64)
			// Attempt to reauthorize the customer
			_, err := h.CustomerAuthRepository.ReauthCustomer(uint(accessLoginID))
			if err != nil {
				return errorhandler.ErrorHandler(c, err, "Customer Not Found", http.StatusInternalServerError)
			}
		} else {
			// If neither adminLogin nor customerLogin exists, return an error
			return errorhandler.ErrorHandler(c, nil, "Admin or Customer Not Found", http.StatusInternalServerError)
		}
	} else {
		// Use adminLogin if it exists
		accessLoginID := accessLogin["id"].(float64)
		_, err := h.AdminAuthRepository.Reauth(uint(accessLoginID))

		if err != nil {
			return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
		}
	}

	creditlimit, err := h.CreditLimitRepository.ListCreditLimit()
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
}
func (h *handlerCreditLimit) CreditLimitByID(c echo.Context) error {
	accessLogin, ok := c.Get("adminLogin").(jwt.MapClaims)
	if !ok || accessLogin == nil {
		userLogin, _ := c.Get("customerLogin").(jwt.MapClaims)
		if userLogin != nil {

			accessLoginID := userLogin["id"].(float64)
			// Attempt to reauthorize the customer
			_, err := h.CustomerAuthRepository.ReauthCustomer(uint(accessLoginID))
			if err != nil {
				return errorhandler.ErrorHandler(c, err, "Customer Not Found", http.StatusInternalServerError)
			}
		} else {
			// If neither adminLogin nor customerLogin exists, return an error
			return errorhandler.ErrorHandler(c, nil, "Admin or Customer Not Found", http.StatusInternalServerError)
		}
	} else {
		// Use adminLogin if it exists
		accessLoginID := accessLogin["id"].(float64)
		_, err := h.AdminAuthRepository.Reauth(uint(accessLoginID))

		if err != nil {
			return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
		}
	}

	ID := c.Param("id")
	LimitID, _ := strconv.Atoi(ID)

	creditlimit, err := h.CreditLimitRepository.CreditLimitByID(uint(LimitID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
}
