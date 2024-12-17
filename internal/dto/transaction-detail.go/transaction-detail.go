package transactiondetaildto

type RequestTransactionDetail struct {
	OTR               string `form:"otr" validate:"required"`
	AdminFee          string `form:"admin_fee" validate:"required"`
	InstallmentAmount string `form:"installment_amount" validate:"required"`
	InterestAmount    string `form:"interest_amount" validate:"required"`
	CreditLimitID     uint   `form:"credit_limit_id" validate:"required"`
	PartnerBankID     uint   `form:"partner_bank_id" validate:"required"`
	AssetID           uint   `form:"asset_id" validate:"required"`
	PartnerID         uint   `form:"partner_id"`
}
