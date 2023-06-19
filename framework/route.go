package capruk

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setupRoute() *echo.Echo {

	r := echo.New()
	r.Use(middleware.Recover())
	r.Use(middleware.CORS())

	return r
}

func GROUP(grup string, m ...echo.MiddlewareFunc) *echo.Group {
	g := appRoute.Group(grup, m...)
	return g
}

func GET(l string, h echo.HandlerFunc) {
	appRoute.GET(l, h)
}

func POST(l string, h echo.HandlerFunc) {
	appRoute.POST(l, h)
}

func DELETE(l string, h echo.HandlerFunc) {
	appRoute.DELETE(l, h)
}

func MID(group *echo.Group, m ...echo.MiddlewareFunc) {
	group.Use(m...)
}

func Mid(m ...echo.MiddlewareFunc) {
	appRoute.Use(m...)
}

func startServer(httport int) {
	addr := fmt.Sprintf(":%d", httport)
	err := appRoute.Start(addr)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
