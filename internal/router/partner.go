package router

import (
	"kredit-plus/internal/handler"
	repositories "kredit-plus/internal/repository"
	"kredit-plus/pkg/middleware"
	"kredit-plus/pkg/mysql"

	"github.com/labstack/echo/v4"
)

func Partner(e *echo.Group) {
	repo := repositories.RepositoryCreditLimit(mysql.DB)
	h := handler.HandlerPartner(repo, repo, repo, repo)
	//only admin
	e.POST("/admin/partner/register", middleware.Auth(h.CreatePartner))
	e.GET("/partner/:id", middleware.Auth(h.PartnerByID))
	e.GET("/admin/partner/list", middleware.Auth(h.ListPartner))
	//admin && customer

}
