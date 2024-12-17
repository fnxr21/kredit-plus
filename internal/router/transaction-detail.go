package router

import (
	"kredit-plus/internal/handler"
	repositories "kredit-plus/internal/repository"
	"kredit-plus/pkg/middleware"
	"kredit-plus/pkg/mysql"

	"github.com/labstack/echo/v4"
)

func TransactionDetail(e *echo.Group) {
	repo := repositories.RepositoryCreditLimit(mysql.DB)
	h := handler.HandlerTransactionDetail(repo, repo, repo)
	//only admin
	e.POST("/transactiondetail/register", middleware.Auth(h.CreateTransactionDetail))
	e.GET("/transactiondetail/:id", middleware.Auth(h.TransactionDetailByID))
	e.GET("/transactiondetail/list", middleware.Auth(h.ListTransactionDetail))
	//admin && customer

}
