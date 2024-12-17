package transactionpaymentdto

type RequestTransactionPayment struct {
	TransactionDetailID uint   `form:"transaction_detail_id"`
	Amount              string `form:"amount"`
	PartnerID           uint   `form:"partner_id"`
}
