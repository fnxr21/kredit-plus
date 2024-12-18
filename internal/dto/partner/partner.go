package partnerdto


type RequestPartner struct {
	Name          string `form:"name"`
	Email         string `form:"email"`
	PhoneNumber   string `form:"phone_number"`
	Address       string `form:"address"`
	PartnerBankID uint   `form:"partner_bank_id"`
}
