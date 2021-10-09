package server

import (
	"github.com/Nikita99100/CarCRUDService/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Rest struct {
	Router  *echo.Echo
	Handler *handler.Handler
}

func (r *Rest) Route() {
	r.Router.GET("/", r.hello)
}

func (r *Rest) hello(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>Hello from car CRUD service</h1>")
}
