package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	webServer := echo.New()
	webServer.HideBanner = true
	webServer.Use(middleware.Logger())
	webServer.Use(middleware.Recover())

	webServer.GET("/health", func(c echo.Context) error {
		return c.JSON(200, echo.Map{"status": "ok"})
	})

	return webServer
}
