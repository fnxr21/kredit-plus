package handler

import (
	"fmt"
	"kredit-plus/internal/models"
	"kredit-plus/pkg/bcrypt"

	"github.com/stretchr/testify/mock"
)

type MockCustomerRepository struct {
	Mock mock.Mock
}

func mockCustomerData() *models.Customer {
	hashedPassword, err := bcrypt.HashingPassword("mypassword")

	if err != nil {
		fmt.Errorf("Error hashing password : %v", err)
	}

	return &models.Customer{
		Username:    "John",
		Password:    hashedPassword,
		Email:       "john@example.com",
		PhoneNumber: "087623671",
		Nik:         "20003990",
		FullName:    "John F kenedy",
		LegalName:   "John F kenedy",
		Birthplace:  "Jakarta",
		BirthDate:   "2002-04-05",
		Salary:      "8000000",
		ImageKTP:    "path/image",
		ImageSelfie: "path/image",
	}
}
