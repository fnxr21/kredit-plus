package router

import (
	"kredit-plus/internal/handler"
	repositories "kredit-plus/internal/repository"
	"kredit-plus/pkg/middleware"
	"kredit-plus/pkg/mysql"

	"github.com/labstack/echo/v4"
)

func AdminAuth(e *echo.Group) {
	repo := repositories.RepositoryAdminAuth(mysql.DB)
	h := handler.HandlerAdminAuth(repo)
	e.POST("/admin/register", h.RegisterAdmin)
	e.POST("/admin/login", h.Login)
	e.GET("/admin/reauth", middleware.Auth(h.ReauthAdmin))
	e.GET("/admin/logout", middleware.Auth(h.LogoutAdmin))
}
