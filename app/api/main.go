package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/danisbagus/shopping-cart-api/interface/api"
	"github.com/danisbagus/shopping-cart-api/utils/config"
	mid "github.com/danisbagus/shopping-cart-api/utils/rest/middleware"
)

func main() {
	// load config
	config.LoadConfig()

	// new echo instant
	e := echo.New()

	// middleware
	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		}),
		middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
			Skipper: mid.APILogSkipper,
			Handler: mid.APILogHandler,
		}),
	)

	// route
	api.API(e)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
