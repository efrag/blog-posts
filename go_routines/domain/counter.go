package domain

import (
	"github.com/efrag/blog-posts/go_routines/logger"
)

type Counter struct {
	Id       int
	registry int64     // how much money we have collected in this counter
	people   []*Person // which people from the queue were processed at this counter
}

func (c *Counter) PeopleIDs() []int {
	ids := make([]int, 0)
	for _, p := range c.people {
		ids = append(ids, p.Id)
	}
	return ids
}

func (c *Counter) Process(p *Person) {
	pc := p.processingCycles()

	for i := 0; i < pc; i++ {
		// do nothing here basically just simulate the time
		// it would take to actually process the items in
		// this person's basket
	}

	c.people = append(c.people, p)
	c.registry += p.totalCost()

	logger.Log("(Person, #Items, Go Routine): %4v %4d %4v\n", p.Id, p.Items, logger.GetRoutineID())
}
