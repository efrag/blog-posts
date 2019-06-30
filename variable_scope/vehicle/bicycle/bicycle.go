package bicycle

import (
	"fmt"
	"math/rand"
)

var numberOfBicycles = 5

type Bicycle struct {
	HasHorn           bool
	ManufacturingYear int
	Name              string
}

func CreateBicycles() []*Bicycle {
	bicycles := make([]*Bicycle, 0)

	for i := 0; i < numberOfBicycles; i++ {
		bicycles = append(bicycles, &Bicycle{
			HasHorn:           true,
			ManufacturingYear: 2000 + i,
			Name:              "MyBicycleYear",
		})
	}

	return bicycles
}

func ShadowingNumberOfBicycles() {
	fmt.Println("\nBeginning of the func:", numberOfBicycles)

	if numberOfBicycles := rand.Intn(100); numberOfBicycles > 0 {
		for i := 0; i < 4; i++ {
			numberOfBicycles += 1
			fmt.Println("-- Inside the for:", numberOfBicycles)
		}
	}

	fmt.Println("End of the func:", numberOfBicycles)
}
