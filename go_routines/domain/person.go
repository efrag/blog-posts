package domain

import "math/rand"

type Person struct {
	Id    int
	Items int
}

// TotalCost calculates the cost of all items
func (p *Person) totalCost() int64 {
	total := 0
	for i := 0; i < p.Items; i++ {
		// so that our range is between 1 to 100 inclusive
		total += rand.Intn(100) + 1
	}
	return int64(total)
}

func (p *Person) processingCycles() int {
	// introducing an artificial delay when processing the
	// items of a specific user. The delay varies based on
	// the number of items
	return 10000000 * p.Items
}
