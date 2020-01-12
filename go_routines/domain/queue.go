package domain

import "sync"

type Queue struct {
	people []*Person
	Lock   *sync.Mutex
}

// Pop will remove the top element from the queue and return it
func (q *Queue) Pop() *Person {
	if len(q.people) > 0 {
		person := &Person{}
		person, q.people = q.people[0], q.people[1:]
		return person
	}

	return nil
}

// Push adds an element at the end of the existing queue
func (q *Queue) Push(person *Person) {
	q.people = append(q.people, person)
}

// NumberOfPeople
//returns the number of people in our queue
func (q *Queue) NumberOfPeople() int {
	return len(q.people)
}

func (q *Queue) Init(p []*Person, m *sync.Mutex) {
	q.people = p
	q.Lock = m
}
