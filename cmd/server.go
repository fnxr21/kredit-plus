package cmd

import (
	// "fmt"
	"kredit-plus/internal/database"
	"kredit-plus/internal/router"
	"kredit-plus/pkg/log"
	"kredit-plus/pkg/mysql"
	"kredit-plus/pkg/validate"
	"net/http"
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

	e.Validator = validate.New()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	
		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}
	log.Init()
	// initAllPkg(e)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))
	//route end point init
	router.RouteInt(e.Group("/api/v1"))

	PORT := os.Getenv("APP_PORT")

	// default port 500
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

func initAllPkg(e *echo.Echo) {
	// connectdatabase
	mysql.DataBaseinit()
	database.RunMigration()

	e.Validator = validate.New()
	log.Init()
}
