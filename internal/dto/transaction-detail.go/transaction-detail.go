package transactiondetaildto

type RequestTransactionDetail struct {
	ContractNumber    uint    `json:"contract_number"`
	OTR               string `json:"otr"`
	AdminFee          string `json:"admin_fee"`
	InstallmentAmount string `json:"installment_amount"`
	InterestAmount    string `json:"interest_amount"`
	Status            string  `json:"payment"`
	CreditLimitID     uint    `json:"credit_limit_id"`
	PartnerBankID     uint    `json:"partner_bank_id"`
	AssetID           uint    `json:"asset_id"`
	PartnerID         uint    `json:"partner_id"`
	CustomerID        uint    `json:"customer_id"`
}
