package proto

type Car struct {
	UUID    string
	Brand   string
	Model   string
	Price   *int
	State   string
	Mileage *int
}

var CarsMap map[string]*Car

func init() {
	CarsMap = make(map[string]*Car)
}

func (c *Car) UpdateNotNull(upd *Car) error {
	if upd.Brand != "" {
		c.Brand = upd.Brand
	}
	if upd.Model != "" {
		c.Model = upd.Model
	}
	if upd.Price != nil {
		c.Price = upd.Price
	}
	if upd.State != "" {
		c.State = upd.State
	}
	if upd.Mileage != nil {
		c.Mileage = upd.Mileage
	}
	return nil
}

func FindCarsByFilter(filter *Car) *[]Car {
	var foundCars []Car
	for _, v := range CarsMap {
		if filter.UUID != "" {
			if v.UUID != filter.UUID {
				continue
			}
		}
		if filter.Model != "" {
			if v.Model != filter.Model {
				continue
			}
		}
		if filter.Brand != "" {
			if v.Brand != filter.Brand {
				continue
			}
		}
		if filter.State != "" {
			if v.State != filter.State {
				continue
			}
		}
		if filter.Price != nil {
			if *v.Price != *filter.Price {
				continue
			}
		}
		if filter.Mileage != nil {
			if *v.Mileage != *filter.Mileage {
				continue
			}
		}
		foundCars = append(foundCars, *v)
	}
	return &foundCars
}

const (
	// Статусы автомобиля

	CarStateOnWay     = "on-way"    //В пути
	CarStateInStock   = "in-stock"  //На складе
	CarStateSold      = "sold"      //продан
	CarStateWithdrawn = "withdrawn" //снят с продажи
)
