package cmd

import (
	"kredit-plus/internal/database"
	"kredit-plus/pkg/mysql"
	"kredit-plus/pkg/validate"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunServer() {

	dotEnv()

	e := echo.New()

	mysql.DataBaseinit()
	database.RunMigration()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))
	e.Validator = validate.New()
	PORT := os.Getenv("APP_PORT")
	// default port
	if PORT == "" {
		PORT = "5000"
	}

	e.Logger.Fatal(e.Start(":" + PORT))
}

func dotEnv() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
}
