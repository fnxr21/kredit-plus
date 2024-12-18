package handler

// import (
// 	"bytes"
// 	"fmt"
// 	authdto "kredit-plus/internal/dto/auth"
// 	"kredit-plus/internal/models"
// 	"kredit-plus/pkg/bcrypt"
// 	"net/http"
// 	"net/http/httptest"
// 	"net/url"
// 	"testing"
// 	"time"

// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockAdminRepository struct {
// 	Mock mock.Mock
// }

// // go test-v // on folder root for check test file exist
// // got tes -v ./internal/handler
// func (m *MockAdminRepository) Login(username string) (models.MyUser, error) {
// 	args := m.Mock.Called(username)
// 	var user models.MyUser
// 	if args.Get(0) != nil {
// 		return args.Get(0).(models.MyUser), args.Error(1)
// 	}
// 	// in empty i should use * here ....
// 	// value must be nil ,err but i forget use *  . . .
// 	return user, args.Error(1)
// }
// func (m *MockAdminRepository) Reauth(id uint) (models.MyUser, error) {
// 	args := m.Mock.Called(id)
// 	var user models.MyUser

// 	if args.Get(0) != nil {
// 		return args.Get(0).(models.MyUser), args.Error(1)
// 	}
// 	return user, args.Error(1)
// }

// func (m *MockAdminRepository) Register(user models.MyUser) (models.MyUser, error) {
// 	args := m.Mock.Called(user)
// 	//this should be only error
// 	return user, args.Error(0)
// }

// // okey lets handle login
// func mockDbUser() models.MyUser {
// 	hashedPassword, err := bcrypt.HashingPassword("mypassword")

// 	if err != nil {
// 		fmt.Errorf("Error hashing password : %v", err)
// 	}

// 	return models.MyUser{
// 		Username:    "fandi",
// 		Password:    hashedPassword,
// 		PhoneNumber: "0842390423",
// 		Email:       "fandi@example.com",
// 	}
// }

// func TestLogin(t *testing.T) {
// 	mockRepo := new(MockAdminRepository)

// 	adminService := HandlerAdminAuth(mockRepo)

// 	t.Run("should return user token", func(t *testing.T) {
// 		payload := models.MyUser{
// 			Email:    "test@example.com",
// 			Password: "mypassword",
// 		}
// 		// Prepare the form data
// 		formData := url.Values{}
// 		formData.Add("email", payload.Email)
// 		formData.Add("password", payload.Password)
// 		// Create the HTTP request with multipart/form-data
// 		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(formData.Encode()))
// 		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 		// Record the HTTP response
// 		rr := httptest.NewRecorder()

// 		// Create Echo instance
// 		e := echo.New()

// 		mockRepo.Mock.On("Login", payload.Username).Return(mockRepo, nil).Once()
// 		router := e.Group("")
// 		router.POST("/login", adminService.Login)
// 		if assert.NoError(t, err) {
// 			assert.Equal(t, http.StatusOK, rr.Code)
// 			assert.JSONEq(t, `{"status":200,"data":"still-active"}`, rec.Body.String())
// 		}
// 	})

// }

// func MapToAuthDTO(user models.MyUser) authdto.LoginRequest {
// 	return authdto.LoginRequest{
// 		Username: user.Username,
// 		Password: user.Password,
// 	}
// }
