package proto

type Car struct {
	UUID    string
	Brand   string
	Model   string
	Price   int
	Status  string
	Mileage int
}

var CarsMap map[string]*Car

func init() {
	CarsMap = make(map[string]*Car)
}
