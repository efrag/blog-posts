package car

type Car struct {
	FuelType          string
	ManufacturingYear int
	Name              string
	NumberOfDoors     int
}

func GetDefaultCar() *Car {
	return &Car{
		FuelType:          "gas",
		ManufacturingYear: 2010,
		Name:              "Audi",
		// fourDoor is an internal pkg constant.
		// It is visible (in scope) because it's
		// defined in the constants.go file and
		// is in the Package scope
		NumberOfDoors: fourDoor,
	}
}
