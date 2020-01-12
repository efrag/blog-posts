package utils

import (
	"github.com/efrag/blog-posts/go_routines/domain"
)

func itemInList(options []string, name string) bool {
	found := false
	for _, option := range options {
		if option == name {
			found = true
			continue
		}
	}
	return found
}

func AllowedOption(options []string, name string) bool {
	return itemInList(options, name)
}

func FlagIsPresent(options []string, name string) bool {
	return itemInList(options, name)
}

func GetDuplicatePeople(counters []*domain.Counter) []int {
	d := make([]int, 0)
	p := map[int]int{}

	for _, c := range counters {
		ids := c.PeopleIDs()
		for i := 0; i < len(ids); i++ {
			p[ids[i]] += 1
		}
	}

	for id, c := range p {
		if c > 1 {
			d = append(d, id)
		}
	}

	return d
}
