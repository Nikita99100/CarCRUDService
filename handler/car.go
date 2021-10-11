package handler

import (
	"github.com/Nikita99100/CarCRUDService/proto"
	guuid "github.com/google/uuid"
	"github.com/pkg/errors"
)

func (h *Handler) CreateCar(car *proto.Car) (*proto.Car, error) {
	//if uuid is empty, then generate
	if car.UUID == "" {
		car.UUID = guuid.New().String()
	} else {
		//if a car with this uuid already exists, then return error
		_, ok := proto.CarsMap[car.UUID]
		if ok {
			return nil, errors.New("car already exists")
		}
	}
	proto.CarsMap[car.UUID] = car
	return car, nil
}

func (h *Handler) GetCarByUUID(uuid string) (*proto.Car, error) {
	car, ok := proto.CarsMap[uuid]
	if !ok {
		return nil, errors.New("car not found")
	}
	return car, nil
}

func (h *Handler) GetFilteredCars(filter *proto.Car) (*[]proto.Car, error) {
	foundCars := proto.FindCarsByFilter(filter)
	return foundCars, nil
}

func (h *Handler) UpdateCar(updates *proto.Car) error {
	car, ok := proto.CarsMap[updates.UUID]
	if !ok {
		return errors.New("car not found")
	}
	if err := car.UpdateNotNull(updates); err != nil {
		return errors.Wrap(err, "failed to update not null")
	}
	return nil
}

func (h *Handler) DeleteCar(uuid string) error {
	_, ok := proto.CarsMap[uuid]
	if !ok {
		return errors.New("car does not exist")
	}
	delete(proto.CarsMap, uuid)
	return nil
}
