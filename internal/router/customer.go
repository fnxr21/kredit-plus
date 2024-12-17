package router

import (
	"kredit-plus/internal/handler"
	repositories "kredit-plus/internal/repository"
	"kredit-plus/pkg/middleware"
	"kredit-plus/pkg/mysql"

	"github.com/labstack/echo/v4"
)

func Customer(e *echo.Group) {
	repo := repositories.RepositoryAdminAuth(mysql.DB)
	h := handler.HandlerCustomer(repo, repo)
	e.GET("/admin/customer/:id", middleware.Auth(h.CustomerByID))
	e.GET("/admin/customer/list", middleware.Auth(h.CustomerByList))
	e.GET("/admin/customer/ktp/:id", middleware.Auth(h.CustomerKTPByID))
	e.GET("/admin/customer/selfie/:id", middleware.Auth(h.CustomerSelfieByID))
}
