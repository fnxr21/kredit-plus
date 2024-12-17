package router

import (
	"kredit-plus/internal/handler"
	repositories "kredit-plus/internal/repository"
	"kredit-plus/pkg/middleware"
	"kredit-plus/pkg/mysql"

	"github.com/labstack/echo/v4"
)

func Asset(e *echo.Group) {
	repo := repositories.RepositoryAsset(mysql.DB)
	h := handler.HandlerAsset(repo, repo, repo)
	e.POST("/asset/register", middleware.Auth(h.CreateAsset))
	e.GET("/asset/:id", middleware.Auth(h.AssetByID))
	e.GET("/asset/list", middleware.Auth(h.ListAsset))
}
