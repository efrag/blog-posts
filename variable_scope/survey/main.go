package main

import (
	"fmt"
	"github.com/efrag/blog-posts/variable_scope/vehicle/bicycle"
	"github.com/efrag/blog-posts/variable_scope/vehicle/car"
)

func main() {
	num := 0

	first := car.GetDefaultCar()

	second := &car.Car{
		FuelType:          "petrol",
		ManufacturingYear: 2000,
		Name:              "Toyta",
		NumberOfDoors:     car.ThreeDoor, // this is exported from the car pkg
	}

	fmt.Printf("Option 1: %+v\n", first)
	fmt.Printf("Option 2: %+v\n", second)
	num += 2

	bicyclesSequence := bicycle.CreateBicycles()
	for i := 1; i <= len(bicyclesSequence); i++ {
		num += 1
		fmt.Printf("Option %d: %+v\n", num, bicyclesSequence[i-1])
	}

	bicycle.ShadowingNumberOfBicycles()
}
