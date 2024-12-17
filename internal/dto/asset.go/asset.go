package assetdto

type RequestRegisterAsset struct {
	Name       string `form:"name"`   // Nama aset
	Type       string `form:"type"`   // Jenis aset
	Amount     string `form:"amount"` // Nilai total aset
	PartnerID  uint   `form:"partner_id"`
	CustomerID uint   `form:"customer_id"`
}
