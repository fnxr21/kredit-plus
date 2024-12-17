package router

import (
	"kredit-plus/internal/handler"
	repositories "kredit-plus/internal/repository"
	"kredit-plus/pkg/middleware"
	"kredit-plus/pkg/mysql"

	"github.com/labstack/echo/v4"
)

func CustomerAuth(e *echo.Group) {
	repo := repositories.RepositoryCustomerAuth(mysql.DB)
	h := handler.HandlerCustomerAuth(repo)
	e.POST("/customer/register", h.RegisterCustomer)
	e.POST("/customer/login", h.LoginCustomer)
	e.GET("/customer/reauth", middleware.Auth(h.ReauthCustomer))
}
