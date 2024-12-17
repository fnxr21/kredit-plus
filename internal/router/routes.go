package router

import "github.com/labstack/echo/v4"

func RouteInt(r *echo.Group) {
	AdminAuth(r)
	CustomerAuth(r)
	CreditLimit(r)
	Customer(r)
	PartnerBank(r)
	Partner(r)
	Asset(r)
	TransactionDetail(r)
	TransactionPayment(r)
}
