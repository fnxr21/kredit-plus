package router

import "github.com/labstack/echo/v4"

func RouterInt(r *echo.Group) {
	AdminAuth(r)
}
