package routes

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	E *echo.Echo
}

type RouterGroup struct {
	G *echo.Group
}

type R struct {
	Router *Router
}

var Route *R

func SetupRouter(e *echo.Echo) {
	var (
		r *Router
	)
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	r = &Router{E: e}
	Route = &R{Router: r}
}

func (r *Router) Get(url string, f echo.HandlerFunc) {
	r.E.GET(url, f)
}

func (r *Router) Post(url string, f echo.HandlerFunc) {
	r.E.POST(url, f)
}

func (r *R) CreateGroup(prefix string) *RouterGroup {
	var (
		grup *RouterGroup
		L    *echo.Group
	)
	L = r.Router.E.Group(prefix)
	grup = &RouterGroup{G: L}
	return grup
}

func (rg *RouterGroup) Midd(m echo.MiddlewareFunc) {
	rg.G.Use(m)
}

func (rg *RouterGroup) Get(url string, f echo.HandlerFunc) {
	rg.G.GET(url, f)
}

func (rg *RouterGroup) Post(url string, f echo.HandlerFunc) {
	rg.G.POST(url, f)
}

func (r *R) Start(sPort string) {
	log.Fatal(r.Router.E.Start(sPort))
}
