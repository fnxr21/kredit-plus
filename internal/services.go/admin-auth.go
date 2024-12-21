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
	isValid := bcrypt.CheckPasswordHash(request.Password, admin.Password)

	if !isValid {
		return "Password Incorect", err
	}

	claims := jwt.MapClaims{}
	claims["id"] = admin.ID
	claims["name"] = admin.Username
	claims["status"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return "Failed Generate", err

	}

	return token, nil

}

func (h *serviceAdminAuth) ReauthAdmin(id uint) (string, error) {

	user, err := h.AdminAuthRepository.Reauth(id)
	if err != nil {

		return "Admin Not Found", err
	}

	return user.Username, nil

}

func (h *serviceAdminAuth) RegisterAdmin(request *admindto.RequestRegisterAdmin) (string, error) {

	pass, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return "Failed Hashing Password", err
	}
	admin := models.MyUser{
		Username:    request.Username,
		Password:    pass,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
	}
	user, err := h.AdminAuthRepository.Register(admin)
	if err != nil {
		return "Failed Register", err
	}
	return user.Username + " " + "Register", nil

}
