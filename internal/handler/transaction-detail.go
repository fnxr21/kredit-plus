package handler

import (
	dto "kredit-plus/internal/dto/result"
	transactiondetaildto "kredit-plus/internal/dto/transaction-detail.go"
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

type handlerTransactionDetail struct {
	TransactionDetailRepository repositories.TransactionDetail
	AdminAuthRepository         repositories.AdminAuth
	CustomerAuthRepository      repositories.CustomerAuth
}

func HandlerTransactionDetail(
	TransactionDetailRepository repositories.TransactionDetail,
	AdminAuthRepository repositories.AdminAuth,
	CustomerAuthRepository repositories.CustomerAuth) *handlerTransactionDetail {
	return &handlerTransactionDetail{
		TransactionDetailRepository: TransactionDetailRepository,
		AdminAuthRepository:         AdminAuthRepository,
		CustomerAuthRepository:      CustomerAuthRepository}
}

func (h *handlerTransactionDetail) CreateTransactionDetail(c echo.Context) error {
	accessLogin := c.Get("customerLogin").(jwt.MapClaims)
	accessLoginID := accessLogin["id"].(float64)

	_, err := h.CustomerAuthRepository.ReauthCustomer(uint(accessLoginID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
	}
	// catch json  type
	request := new(transactiondetaildto.RequestTransactionDetail)
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
	OTR, _ := strconv.ParseFloat(strings.TrimSpace(request.OTR), 64)
	AdminFee, _ := strconv.ParseFloat(strings.TrimSpace(request.AdminFee), 64)
	InstallmentAmount, _ := strconv.ParseFloat(strings.TrimSpace(request.InstallmentAmount), 64)
	InterestAmount, _ := strconv.ParseFloat(strings.TrimSpace(request.InterestAmount), 64)

	now := time.Now()

	// Format the time to "YYYYMMDDHHMMSS"
	contractNumberStr := now.Format("20060102150405")

	detail := models.TransactionDetail{
		//
		ContractNumber:    contractNumberStr,
		OTR:               OTR,
		AdminFee:          AdminFee,
		InstallmentAmount: InstallmentAmount,
		InterestAmount:    InterestAmount,
		Status:            "pending",
		CreditLimitID:     request.CreditLimitID,
		PartnerBankID:     request.PartnerBankID,
		AssetID:           request.AssetID,
		PartnerID:         request.PartnerID,
		CustomerID:        uint(accessLoginID),
	}

	res, err := h.TransactionDetailRepository.CreateTransactionDetail(detail)
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Failed Create Limit", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: res})
}

func (h *handlerTransactionDetail) ListTransactionDetail(c echo.Context) error {
	accessLogin := c.Get("adminLogin").(jwt.MapClaims)
	accessLoginID := accessLogin["id"].(float64)

	_, err := h.AdminAuthRepository.Reauth(uint(accessLoginID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
	}

	creditlimit, err := h.TransactionDetailRepository.ListTransactionDetail()
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
}
func (h *handlerTransactionDetail) TransactionDetailByID(c echo.Context) error {
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

	creditlimit, err := h.TransactionDetailRepository.TransactionDetailByID(uint(LimitID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
}
