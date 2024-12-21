package router

import (
	// authdto "kredit-plus/internal/dto/auth"
	// dto "kredit-plus/internal/dto/result"
	// "kredit-plus/internal/handler"
	"kredit-plus/internal/handler"
	repositories "kredit-plus/internal/repository"
	service "kredit-plus/internal/services.go"

	// errorhandler "kredit-plus/pkg/error"
	"kredit-plus/pkg/middleware"
	"kredit-plus/pkg/mysql"

	// "net/http"

	"github.com/labstack/echo/v4"
)

func AdminAuth(e *echo.Group) {
	repo := repositories.RepositoryAdminAuth(mysql.DB)
	service := service.ServiceAdminAuth(repo)
	h := handler.HandlerAdminAuthtest(service)

	e.GET("/admin/login", h.LoginAdmin)
	e.POST("/admin/register", h.RegisterAdmin)
	e.GET("/admin/reauth", middleware.Auth(h.ReauthAdmin))
	e.GET("/admin/logout", middleware.Auth(h.LogoutAdmin))
}
