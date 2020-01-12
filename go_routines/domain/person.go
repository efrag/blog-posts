package domain

import "math/rand"

type Person struct {
	Id    int
	Items int
}

func (p *Person) totalCost() int64 {
	total := 0
	for i := 0; i < p.Items; i++ {
		total += rand.Intn(100) + 1 // so that our range is between 1 to 100 inclusive
	}
	return int64(total)
}

func (p *Person) processingCycles() int {
	// this just for introducing an artificial delay when processing the items for a
	// specific user. We want the delay to vary based on the number of items
	return 10000000 * p.Items
}
