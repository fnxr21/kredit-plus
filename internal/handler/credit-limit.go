package handler

import (
	creditlimitdto "kredit-plus/internal/dto/credit-limit"
	dto "kredit-plus/internal/dto/result"
	"kredit-plus/internal/models"
	repositories "kredit-plus/internal/repository"
	errorhandler "kredit-plus/pkg/error"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerCreditLimit struct {
	CreditLimitRepository repositories.CreditLimit
	AdminAuthRepository   repositories.AdminAuth
}

func HandlerCreditLimit(CreditLimitRepository repositories.CreditLimit, AdminAuthRepository repositories.AdminAuth) *handlerCreditLimit {
	return &handlerCreditLimit{CreditLimitRepository: CreditLimitRepository, AdminAuthRepository: AdminAuthRepository}
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
