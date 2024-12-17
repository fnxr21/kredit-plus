package creditlimitdto

type RequestRegisterCustomer struct {
	TenorMonths string `form:"tenor_months" validate:"required"`
	LimitAmount string `form:"limit_amount" validate:"required"`
}
