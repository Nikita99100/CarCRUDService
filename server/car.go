package server

import (
	"github.com/Nikita99100/CarCRUDService/proto"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

func (r *Rest) CreateCar(c echo.Context) error {
	var err error
	car := new(proto.Car)
	if err = c.Bind(car); err != nil {
		return c.JSON(http.StatusBadRequest, "bind error on create car")
	}
	car, err = r.Handler.CreateCar(car)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.Wrap(err, "failed to create car").Error())
	}
	return c.JSON(http.StatusCreated, car)
}

func (r *Rest) GetCarByUUID(c echo.Context) error {
	uuid := c.Param("uuid")
	car, err := r.Handler.GetCarByUUID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.Wrap(err, "failed to get car by uuid").Error())
	}
	return c.JSON(http.StatusOK, car)
}

func (r *Rest) GetFilteredCars(c echo.Context) error {
	filter := new(proto.Car)
	if err := c.Bind(filter); err != nil {
		return c.JSON(http.StatusBadRequest, "bind error on get filtered car")
	}
	cars, err := r.Handler.GetFilteredCars(filter)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.Wrap(err, "failed to get filtered cars").Error())
	}
	return c.JSON(http.StatusOK, cars)
}

func (r *Rest) UpdateCar(c echo.Context) error {
	updates := new(proto.Car)
	if err := c.Bind(updates); err != nil {
		c.JSON(http.StatusBadRequest, "bind error on update car")
	}
	if err := r.Handler.UpdateCar(updates); err != nil {
		return c.JSON(http.StatusBadRequest, errors.Wrap(err, "failed to handle update car").Error())
	}
	return c.JSON(http.StatusOK, "updated")
}

func (r *Rest) DeleteCar(c echo.Context) error {
	uuid := c.Param("uuid")
	if err := r.Handler.DeleteCar(uuid); err != nil {
		return c.JSON(http.StatusBadRequest, errors.Wrap(err, "failed to delete car").Error())
	}
	return c.JSON(http.StatusOK, "deleted")
}
