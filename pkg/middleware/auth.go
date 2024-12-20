package middleware

import (
	// "log"

	dto "kredit-plus/internal/dto/result"
	jwtToken "kredit-plus/pkg/jwt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
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

		switch status {
		case "customer":
			log.Println("customer", claims)
			c.Set("customerLogin", claims)
			return next(c)
		case "admin":
			log.Println("admin", claims)
			c.Set("adminLogin", claims)
			return next(c)
		default:
			// Optional: handle cases where status doesn't match "customer" or "admin"
			log.Println("unauthorized access", claims)
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		}

	}
}

func GetUserIdFromContext(c echo.Context) int {
	if c == nil {
		return -1
	}

	// Check admin login
	if accessLogin, ok := c.Get("adminLogin").(jwt.MapClaims); ok {
		if id, ok := accessLogin["id"].(float64); ok {
			return int(id)
		}
	}

	// Check customer login
	if userLogin, ok := c.Get("customerLogin").(jwt.MapClaims); ok {
		if id, ok := userLogin["id"].(float64); ok {
			return int(id)
		}
	}

	return -1 // Return -1 if no valid login found
}
