package domain

import (
	"sync"
	"testing"
)

var people = []*Person{
	{
		Id:    0,
		Items: 1,
	},
	{
		Id:    1,
		Items: 2,
	},
}

func TestQueue_Push(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name           string
		person         *Person
		expectedPeople int
	}{
		{"Add 1st person", &Person{Id: 2, Items: 3}, 3},
		{"Add 2nd person", &Person{Id: 3, Items: 4}, 4},
	}

	q := Queue{}
	q.Init(people, &sync.Mutex{})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			q.Push(tc.person)

			np := q.NumberOfPeople()
			if np != tc.expectedPeople {
				t.Errorf("Expecting %v Items but got: %v", tc.expectedPeople, np)
			}
		})
	}
}

func TestQueue_Pop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		exID     int
		exItems  int
		exPeople int
	}{
		{name: "Pop 1st person", exID: 0, exItems: 1, exPeople: 1},
		{name: "Pop 2nd person", exID: 1, exItems: 2, exPeople: 0},
	}

	q := Queue{}
	q.Init(people, &sync.Mutex{})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := q.Pop()
			if p.Id != tc.exID {
				t.Errorf("Invalid Person ID: %v", p.Id)
			}
			if p.Items != tc.exItems {
				t.Errorf("Invalid Person Items: %v", p.Items)
			}
			if q.NumberOfPeople() != tc.exPeople {
				t.Errorf("Invalid number of people in the queue: %v", q.NumberOfPeople())
			}
		})
	}
}
