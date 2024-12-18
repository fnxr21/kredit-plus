package handler

import (
	dto "kredit-plus/internal/dto/result"
	transactionpaymentdto "kredit-plus/internal/dto/transaction-payment"
	"kredit-plus/internal/models"
	repositories "kredit-plus/internal/repository"
	errorhandler "kredit-plus/pkg/error"
	"strconv"
	"strings"
	"time"

	// "log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerTransactionPayment struct {
	TransactionPaymentRepository repositories.TransactionPayment
	AdminAuthRepository          repositories.AdminAuth
	CustomerAuthRepository       repositories.CustomerAuth
}

func HandlerTransactionPayment(
	TransactionPaymentRepository repositories.TransactionPayment,
	AdminAuthRepository repositories.AdminAuth,
	CustomerAuthRepository repositories.CustomerAuth) *handlerTransactionPayment {
	return &handlerTransactionPayment{
		TransactionPaymentRepository: TransactionPaymentRepository,
		AdminAuthRepository:          AdminAuthRepository,
		CustomerAuthRepository:       CustomerAuthRepository}
}

func (h *handlerTransactionPayment) CreateTransactionPayment(c echo.Context) error {
	accessLogin := c.Get("customerLogin").(jwt.MapClaims)
	accessLoginID := accessLogin["id"].(float64)

	_, err := h.CustomerAuthRepository.ReauthCustomer(uint(accessLoginID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
	}
	// catch json  type
	request := new(transactionpaymentdto.RequestTransactionPayment)
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
	Amount, _ := strconv.ParseFloat(strings.TrimSpace(request.Amount), 64)

	// Format the time to "YYYYMMDDHHMMSS"

	detail := models.TransactionPayment{
		TransactionDetailID: 1,
		Status:              "pending",
		Amount:              Amount,
		PaymentDate:         time.Now(),
		PartnerID:           request.PartnerID,

		CustomerID: uint(accessLoginID),
	}

	res, err := h.TransactionPaymentRepository.CreateTransactionPayment(detail)
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Failed Create Limit", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: res})
}

func (h *handlerTransactionPayment) ListTransactionPayment(c echo.Context) error {
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

	creditlimit, err := h.TransactionPaymentRepository.ListTransactionPayment()
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
}
func (h *handlerTransactionPayment) TransactionPaymentByID(c echo.Context) error {
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

	creditlimit, err := h.TransactionPaymentRepository.TransactionPaymentByID(uint(LimitID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
}
