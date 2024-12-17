package handler

// import (
// 	"kredit-plus/internal/models"

// 	"github.com/stretchr/testify/mock"
// )

// type MockAuthRepository struct {
// 	Mock mock.Mock
// }

// func (m *MockAuthRepository) Register(admin models.MyUser) (models.MyUser, error) {
// 	args := m.Mock.Called(admin)
// 	return admin, args.Error(0)
// }

// func (m *MockAuthRepository) Login(username string) (models.MyUser, error) {
// 	args := m.Mock.Called(username)
// 	admin := models.MyUser{}
// 	if args.Get(0) != nil {
// 		return args.Get(0).(models.MyUser), args.Error(1)
// 	}
// 	return admin, args.Error(1)
// }
