package handler

import (
	partnerdto "kredit-plus/internal/dto/partner"
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

type handlerPartner struct {
	PartnerRepository      repositories.Partner
	PartnerBankRepository  repositories.PartnerBank
	AdminAuthRepository    repositories.AdminAuth
	CustomerAuthRepository repositories.CustomerAuth
}

func HandlerPartner(
	PartnerRepository repositories.Partner,
	AdminAuthRepository repositories.AdminAuth,
	PartnerBankRepository repositories.PartnerBank,
	CustomerAuthRepository repositories.CustomerAuth) *handlerPartner {
	return &handlerPartner{
		PartnerRepository:      PartnerRepository,
		PartnerBankRepository:  PartnerBankRepository,
		AdminAuthRepository:    AdminAuthRepository,
		CustomerAuthRepository: CustomerAuthRepository}
}

func (h *handlerPartner) CreatePartner(c echo.Context) error {
	accessLogin := c.Get("adminLogin").(jwt.MapClaims)
	accessLoginID := accessLogin["id"].(float64)

	_, err := h.AdminAuthRepository.Reauth(uint(accessLoginID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
	}
	// catch json  type
	request := new(partnerdto.RequestPartner)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	
	error := c.Validate(request)

	if error != nil {
		return errorhandler.ErrorHandler(c, error, error.Error(), http.StatusBadRequest)
	}

	//check if bank exist
	bankID, err := h.PartnerBankRepository.PartnerBankByID(request.PartnerBankID)
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Bank Not Found", http.StatusBadRequest)
	}

	partner := models.Partner{
		Name:          request.Name,
		Email:         request.Email,
		PhoneNumber:   request.PhoneNumber,
		Address:       request.Address,
		PartnerBankID: bankID.ID,
	}

	creditlimit, err := h.PartnerRepository.CreatePartner(partner)
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Failed Create Limit", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
}

func (h *handlerPartner) ListPartner(c echo.Context) error {
	accessLogin := c.Get("adminLogin").(jwt.MapClaims)
	accessLoginID := accessLogin["id"].(float64)

	_, err := h.AdminAuthRepository.Reauth(uint(accessLoginID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
	}

	creditlimit, err := h.PartnerRepository.ListPartner()
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
}
func (h *handlerPartner) PartnerByID(c echo.Context) error {
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

	creditlimit, err := h.PartnerRepository.PartnerByID(uint(LimitID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
}
