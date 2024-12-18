package router

import (
	"kredit-plus/internal/handler"
	repositories "kredit-plus/internal/repository"
	"kredit-plus/pkg/middleware"
	"kredit-plus/pkg/mysql"

	"github.com/labstack/echo/v4"
)

func TransactionPayment(e *echo.Group) {
	repo := repositories.RepositoryCreditLimit(mysql.DB)
	h := handler.HandlerTransactionPayment(repo, repo, repo)
	//only admin
	e.POST("/transactionpayment/register", middleware.Auth(h.CreateTransactionPayment))
	e.GET("/transactionpayment/:id", middleware.Auth(h.TransactionPaymentByID))
	e.GET("/transactionpayment/list", middleware.Auth(h.ListTransactionPayment))
	//admin && customer

}
