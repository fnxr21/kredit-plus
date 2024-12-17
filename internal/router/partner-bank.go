package router

import (
	"kredit-plus/internal/handler"
	repositories "kredit-plus/internal/repository"
	"kredit-plus/pkg/middleware"
	"kredit-plus/pkg/mysql"

	"github.com/labstack/echo/v4"
)

func PartnerBank(e *echo.Group) {
	repo := repositories.RepositoryPartner(mysql.DB)
	h := handler.HandlerPartnerBank(repo, repo, repo)
	//only admin
	e.POST("/bank/register", middleware.Auth(h.CreatePartnerBank))
	e.GET("/bank/list", middleware.Auth(h.ListPartnerBank))
	e.GET("/bank/:id", middleware.Auth(h.PartnerBankByID))

}
