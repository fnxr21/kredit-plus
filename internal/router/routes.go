package router

import "github.com/labstack/echo/v4"

func RouteInt(r *echo.Group) {
	AdminAuth(r)
	CustomerAuth(r)
}
