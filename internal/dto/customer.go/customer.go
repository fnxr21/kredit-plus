package customerdto

type RequestRegisterCustomer struct {
	Username    string `form:"username" validate:"required"`
	Password    string `form:"password" validate:"required"`
	Email       string `form:"email" validate:"required"`
	PhoneNumber string `form:"phone_number" validate:"required"`
	Nik         string `form:"nik" validate:"required"`
	FullName    string `form:"full_name" validate:"required"`
	LegalName   string `form:"legal_name" validate:"required"`
	Birthplace  string `form:"birthplace" validate:"required"`
	BirthDate   string `form:"dob" validate:"required"`
	Salary      string `form:"salary" validate:"required"`
}

//cancel ... to risky for shared
type ResponseRegisterCustomer struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
	Birthplace  string `json:"birthplace"`
	BirthDate   string `json:"dob"`
	Salary      string `json:"salary"`
}
