package handler

import (
	customerdto "kredit-plus/internal/dto/customer.go"
	dto "kredit-plus/internal/dto/result"
	repositories "kredit-plus/internal/repository"
	errorhandler "kredit-plus/pkg/error"
	// handlerimage "kredit-plus/pkg/image"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerCustomer struct {
	CustomerRepository  repositories.Customer
	AdminAuthRepository repositories.AdminAuth
}

func HandlerCustomer(CustomerRepository repositories.Customer, AdminAuthRepository repositories.AdminAuth) *handlerCustomer {
	return &handlerCustomer{CustomerRepository: CustomerRepository, AdminAuthRepository: AdminAuthRepository}
}

func (h *handlerCustomer) CustomerByID(c echo.Context) error {
	accessLogin := c.Get("adminLogin").(jwt.MapClaims)
	accessLoginID := accessLogin["id"].(float64)
	_, err := h.AdminAuthRepository.Reauth(uint(accessLoginID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
	}

	ID := c.Param("id")
	CustomerID, _ := strconv.Atoi(ID)
	user, err := h.CustomerRepository.CustomerByID(CustomerID)
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Customer Not Found", http.StatusUnauthorized)
	}

	//use string for json (image)
	// selfieBase64, err := handlerimage.FileToBase64(user.ImageSelfie)
	// if err != nil {

	// 	return errorhandler.ErrorHandler(c, err, "Error retrieving Selfie", http.StatusInternalServerError)

	// }
	// ktpBase64, err := handlerimage.FileToBase64(user.ImageKTP)
	// if err != nil {
	// 	return errorhandler.ErrorHandler(c, err, "Error retrieving ktp", http.StatusInternalServerError)

	// }
	salary := strconv.FormatFloat(user.Salary, 'f', 2, 64) // 2 decimal places
	response := customerdto.ResponseCustomerID{
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Nik:         user.Nik,
		FullName:    user.FullName,
		LegalName:   user.LegalName,
		Birthplace:  user.Birthplace,
		BirthDate:   user.BirthDate,
		Salary:      salary,
		// ImageKTP:    ktpBase64,
		// ImageSelfie: selfieBase64,
	}
	return c.JSON(http.StatusOK, dto.SuccessReauth{Status: http.StatusOK, Data: response})

}
func (h *handlerCustomer) CustomerKTPByID(c echo.Context) error {
	accessLogin := c.Get("adminLogin").(jwt.MapClaims)
	accessLoginID := accessLogin["id"].(float64)
	_, err := h.AdminAuthRepository.Reauth(uint(accessLoginID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
	}

	ID := c.Param("id")
	CustomerID, _ := strconv.Atoi(ID)
	user, err := h.CustomerRepository.CustomerByID(CustomerID)
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Customer Not Found", http.StatusUnauthorized)
	}

	return c.File(user.ImageKTP)

}
func (h *handlerCustomer) CustomerSelfieByID(c echo.Context) error {
	accessLogin := c.Get("adminLogin").(jwt.MapClaims)
	accessLoginID := accessLogin["id"].(float64)
	_, err := h.AdminAuthRepository.Reauth(uint(accessLoginID))
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
	}

	ID := c.Param("id")
	CustomerID, _ := strconv.Atoi(ID)
	user, err := h.CustomerRepository.CustomerByID(CustomerID)
	if err != nil {
		return errorhandler.ErrorHandler(c, err, "Customer Not Found", http.StatusUnauthorized)
	}

	return c.File(user.ImageSelfie)

}
