package handler

import (
	partnerbankdto "kredit-plus/internal/dto/partner-bank.go"
	dto "kredit-plus/internal/dto/result"
	"kredit-plus/internal/models"
	repositories "kredit-plus/internal/repository"
	errorhandler "kredit-plus/pkg/error"
	"strconv"

	// "log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerPartnerBank struct {
	PartnerBankRepository  repositories.PartnerBank
	AdminAuthRepository    repositories.AdminAuth
	CustomerAuthRepository repositories.CustomerAuth
}

func HandlerPartnerBank(
	PartnerBankRepository repositories.PartnerBank,
	AdminAuthRepository repositories.AdminAuth,
	CustomerAuthRepository repositories.CustomerAuth) *handlerPartnerBank {
	return &handlerPartnerBank{
		PartnerBankRepository:  PartnerBankRepository,
		AdminAuthRepository:    AdminAuthRepository,
		CustomerAuthRepository: CustomerAuthRepository}
}

func (h *handlerPartnerBank) CreatePartnerBank(c echo.Context) error {
	accessLogin, ok := c.Get("adminLogin").(jwt.MapClaims)
	// var accessLoginStatus string
	if !ok || accessLogin == nil {
		userLogin, _ := c.Get("customerLogin").(jwt.MapClaims)
		if userLogin != nil {

			accessLoginID := userLogin["id"].(float64)
			// accessLoginStatus = userLogin["status"].(string)
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
		// accessLoginStatus = accessLogin["status"].(string)

		_, err := h.AdminAuthRepository.Reauth(uint(accessLoginID))

		if err != nil {
			return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
		}
	}

	request := new(partnerbankdto.RequestRegisterbank)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	
	error := c.Validate(request)

	if error != nil {
		return errorhandler.ErrorHandler(c, error, error.Error(), http.StatusBadRequest)
	}

	bank := models.PartnerBank{
		BankAccount:       request.BankAccount,
		AccountHolderName: request.AccountHolderName,
		BankName:          request.BankName,
	}
	bank, err := h.PartnerBankRepository.CreatePartnerBank(bank)
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Failed Create Limit", http.StatusBadRequest)
	}
	response := partnerbankdto.ResponseRegisterbank{
		ID:                bank.ID,
		BankAccount:       bank.BankAccount,
		AccountHolderName: bank.AccountHolderName,
		BankName:          bank.BankName,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: response})
}

func (h *handlerPartnerBank) ListPartnerBank(c echo.Context) error {
	accessLogin := c.Get("adminLogin").(jwt.MapClaims)
	// var accessLoginID float64
	accessLoginID := accessLogin["id"].(float64)
	_, err := h.AdminAuthRepository.Reauth(uint(accessLoginID))

	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
	}

	bank, err := h.PartnerBankRepository.ListPartnerBank()
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: bank})
}

func (h *handlerPartnerBank) PartnerBankByID(c echo.Context) error {
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

	creditlimit, err := h.PartnerBankRepository.PartnerBankByID(uint(LimitID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
}
