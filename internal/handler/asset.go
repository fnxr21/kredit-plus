package handler

import (
	assetdto "kredit-plus/internal/dto/asset.go"
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

type handlerAsset struct {
	AssetRepository        repositories.Asset
	AdminAuthRepository    repositories.AdminAuth
	CustomerAuthRepository repositories.CustomerAuth
}

func HandlerAsset(
	AssetRepository repositories.Asset,
	AdminAuthRepository repositories.AdminAuth,
	CustomerAuthRepository repositories.CustomerAuth) *handlerAsset {
	return &handlerAsset{
		AssetRepository:        AssetRepository,
		AdminAuthRepository:    AdminAuthRepository,
		CustomerAuthRepository: CustomerAuthRepository}
}

func (h *handlerAsset) CreateAsset(c echo.Context) error {
	userLogin, _ := c.Get("customerLogin").(jwt.MapClaims)
	accessLoginID := userLogin["id"].(float64)
	// Attempt to reauthorize the customer
	_, err := h.CustomerAuthRepository.ReauthCustomer(uint(accessLoginID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Customer Not Found", http.StatusInternalServerError)
	}
	// catch json  type
	request := new(assetdto.RequestRegisterAsset)
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

	amount, _ := strconv.ParseFloat(strings.TrimSpace(request.Amount), 64)
	limit := models.Asset{
		Name:       request.Name,
		Type:       request.Type,
		Amount:     amount,
		PartnerID:  request.PartnerID,
		CustomerID: uint(accessLoginID),
	}

	creditlimit, err := h.AssetRepository.CreateAsset(limit)
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Failed Create Asset", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
}

func (h *handlerAsset) ListAsset(c echo.Context) error {
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

	creditlimit, err := h.AssetRepository.ListAsset()
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
}

func (h *handlerAsset) AssetByID(c echo.Context) error {
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

	creditlimit, err := h.AssetRepository.AssetByID(uint(LimitID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
}
