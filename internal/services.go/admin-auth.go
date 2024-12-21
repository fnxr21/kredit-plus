package service

import (
	admindto "kredit-plus/internal/dto/admin"
	authdto "kredit-plus/internal/dto/auth"
	// dto "kredit-plus/internal/dto/result"
	"kredit-plus/internal/models"
	repositories "kredit-plus/internal/repository"
	"kredit-plus/pkg/bcrypt"
	// errorservice "kredit-plus/pkg/error"
	jwtToken "kredit-plus/pkg/jwt"
	"log"
	// "net/http"

	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	LoginAdmin(request *authdto.LoginRequest) (string, error)
	ReauthAdmin(id uint) (string, error)
	RegisterAdmin(request *admindto.RequestRegisterAdmin) (string, error)
}

type serviceAdminAuth struct {
	AdminAuthRepository repositories.AdminAuth
}

func ServiceAdminAuth(AdminAuthRepository repositories.AdminAuth) *serviceAdminAuth {
	return &serviceAdminAuth{AdminAuthRepository}
}

func (h *serviceAdminAuth) LoginAdmin(request *authdto.LoginRequest) (string, error) {

	admin, err := h.AdminAuthRepository.Login(request.Username)
	if err != nil {
		return "Admin Not Found", err
	}
	// Compare the provided password with the stored password hash using bcrypt.
	isValid := bcrypt.CheckPasswordHash(request.Password, admin.Password)

	if !isValid {
		return "", err
		// return errorservice.ErrorHandler(c, err, "Incorrect Password", http.StatusBadRequest)
	}

	//generate a JWT token with the user's claims.
	claims := jwt.MapClaims{}
	claims["id"] = admin.ID
	claims["name"] = admin.Username
	claims["status"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // Set token expiration to 2 hours from now.

	//Generate the JWT token using the claims.
	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return "", err

		// return echo.NewHTTPError(http.StatusUnauthorized)
	}
	// response := authdto.LoginResponse{
	// 	Token: token,
	// }

	return token, nil

	// return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: response})
}

// func (h *serviceAdminAuth) ReauthAdmin(c echo.Context) error {
func (h *serviceAdminAuth) ReauthAdmin(id uint) (string, error) {
	// adminLogin := c.Get("adminLogin")
	// adminID := adminLogin.(jwt.MapClaims)["id"].(float64)

	user, err := h.AdminAuthRepository.Reauth(id)
	if err != nil {
		return "", err
		// return errorservice.ErrorHandler(c, err, "Admin Not Found", http.StatusUnauthorized)
	}

	return user.Username, nil
	// return c.JSON(http.StatusOK, dto.SuccessReauth{Status: http.StatusOK, Data: user.Username + " " + "Still Active"})

}

func (h *serviceAdminAuth) RegisterAdmin(request *admindto.RequestRegisterAdmin) (string, error) {
	// request := new(admindto.RequestRegisterAdmin)

	// if err := c.Bind(request); err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	// }

	// error := c.Validate(request)

	// if error != nil {
	// 	return errorservice.ErrorHandler(c, error, error.Error(), http.StatusBadRequest)
	// }
	pass, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return "", err
		// return errorservice.ErrorHandler(c, error, error.Error(), http.StatusBadRequest)
	}
	admin := models.MyUser{
		Username:    request.Username,
		Password:    pass,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
	}
	user, err := h.AdminAuthRepository.Register(admin)
	if err != nil {
		return "", err
		// return errorservice.ErrorHandler(c, err, "Register Failed", http.StatusUnauthorized)
	}
	return user.Username + " " + "Register", nil
	// return c.JSON(http.StatusOK, dto.SuccessReauth{Status: http.StatusOK, Data: user.Username + " " + "Register"})

}

