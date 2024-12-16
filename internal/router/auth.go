package router

import (
	repositories "kredit-plus/internal/repository"
	"kredit-plus/internal/service"
	"kredit-plus/pkg/mysql"

	"github.com/labstack/echo/v4"
)

func AdminAuth(e *echo.Group) {
	repo := repositories.RepositoryAuth(mysql.DB)
	h := service.HandlerAuth(repo)
	e.POST("/admin/login", h.Login)
}
