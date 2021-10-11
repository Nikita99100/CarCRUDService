package server

import (
	"github.com/Nikita99100/CarCRUDService/handler"
	"github.com/labstack/echo/v4"
)

type Rest struct {
	Router  *echo.Echo
	Handler *handler.Handler
}

func (r *Rest) Route() {
	r.Router.POST("/car/create", r.CreateCar)
	r.Router.GET("/car/:uuid", r.GetCarByUUID)
	r.Router.GET("/car", r.GetFilteredCars)
	r.Router.PUT("/car/update", r.UpdateCar)
	r.Router.DELETE("/car/delete/:uuid", r.DeleteCar)
}
