package middleware

import (
	// "log"

	dto "kredit-plus/internal/dto/result"
	jwtToken "kredit-plus/pkg/jwt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// Declare Result struct here ...
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Create Auth function here ...
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusBadRequest, Message: "unauthorized-token"})
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusUnauthorized, Message: "unauthorized"})
		}
		status, _ := claims["status"].(string)
		if status == "customer" {
			log.Println("customer", claims)
			c.Set("customerLogin", claims)
			return next(c)
		}

		if status == "admin" {
			log.Println("admin", claims)
			c.Set("adminLogin", claims)
			return next(c)
		}
		// }

		return c.JSON(http.StatusUnauthorized, "unauthorized-not-found")
	}
}
