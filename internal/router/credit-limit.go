package router

import (
	"kredit-plus/internal/handler"
	repositories "kredit-plus/internal/repository"
	"kredit-plus/pkg/middleware"
	"kredit-plus/pkg/mysql"

	"github.com/labstack/echo/v4"
)

func CreditLimit(e *echo.Group) {
	repo := repositories.RepositoryCreditLimit(mysql.DB)
	h := handler.HandlerCreditLimit(repo, repo, repo)
	//only admin
	e.POST("/admin/credit/register", middleware.Auth(h.CreateCreditLimit))
	//admin && customer
	e.GET("/credit/list", middleware.Auth(h.ListCreditLimit))
	e.GET("/credit/:id", middleware.Auth(h.CreditLimitByID))

}
