package router

import (
	authdto "kredit-plus/internal/dto/auth"
	dto "kredit-plus/internal/dto/result"
	"kredit-plus/internal/handler"
	repositories "kredit-plus/internal/repository"
	errorhandler "kredit-plus/pkg/error"
	"kredit-plus/pkg/middleware"
	"kredit-plus/pkg/mysql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminAuth(e *echo.Group) {
	repo := repositories.RepositoryAdminAuth(mysql.DB)
	h := handler.HandlerAdminAuth(repo)

	repos1 := handler.HandlerAdminAuthtest(repo)
	s := RouterAdminAuthstest(repos1)

	e.POST("/admin/register", h.RegisterAdmin)
	e.POST("/admin/login", h.Login)
	e.GET("/admin/reauth", middleware.Auth(h.ReauthAdmin))
	e.GET("/admin/logout", middleware.Auth(h.LogoutAdmin))
	e.GET("/admin/logintest", s.logins)
}

type routerAdminAuthtest struct {
	AdminAuthhandlertest handler.Authtest
}

func RouterAdminAuthstest(AdminAuthhandlertest handler.Authtest) *routerAdminAuthtest {
	return &routerAdminAuthtest{AdminAuthhandlertest}
}

func (r *routerAdminAuthtest) logins(c echo.Context) error {

	request := new(authdto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	error := c.Validate(request)

	token, err := r.AdminAuthhandlertest.Logins(request)

	if err != nil {
		return errorhandler.ErrorHandler(c, err, error.Error(), http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: token})
}
