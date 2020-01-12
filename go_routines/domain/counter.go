package domain

import (
	"github.com/efrag/blog-posts/go_routines/logger"
)

type Counter struct {
	Id       int
	registry int64     // sum of money collected
	people   []*Person // people who used the counter
}

func (c *Counter) Process(p *Person) {
	pc := p.processingCycles()
	for i := 0; i < pc; i++ {
		// do nothing here - just simulate the time it
		// would take to actually process the items in
		// this person's basket
	}

	c.people = append(c.people, p)
	c.registry += p.totalCost()

	format := "(Person, #Items, Go Routine): %4v %4d %4v\n"
	logger.Log(format, p.Id, p.Items, logger.GetRoutineID())
}

func (c *Counter) PeopleIDs() []int {
	ids := make([]int, 0)
	for _, p := range c.people {
		ids = append(ids, p.Id)
	}
	return ids
}
