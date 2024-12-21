package admindto

type RequestRegisterAdmin struct {
	Username    string `form:"username" validate:"required"`
	Password    string `form:"password" validate:"required"`
	PhoneNumber string `form:"phone_number" validate:"required"`
	Email       string `form:"email" validate:"required,email"`
}
