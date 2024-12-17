package partnerbankdto

type RequestRegisterbank struct {
	BankAccount       string `form:"bank_account" validate:"required"`
	AccountHolderName string `form:"account_holder_name" validate:"required"`
	BankName          string `form:"bank_name" validate:"required"`
}
type ResponseRegisterbank struct {
	ID                uint   `json:"id"`
	BankAccount       string `json:"bank_account"`
	AccountHolderName string `json:"account_holder_name"`
	BankName          string `json:"bank_name"`
}
