package utils

import (
	"context"
	"math/rand"
	"runtime/trace"
	"sync"

	"github.com/efrag/blog-posts/go_routines/domain"
)

func createPeople(number int) []*domain.Person {
	people := make([]*domain.Person, 0)

	for i := 0; i < number; i++ {
		people = append(people, &domain.Person{
			Id:    i,
			Items: rand.Intn(10) + 1, // each person has between 1 and 10 items
		})
	}

	return people
}

func InitializeQueue(ctx context.Context, numPeople int) *domain.Queue {
	ctx, task := trace.NewTask(ctx, "make queue")
	q := &domain.Queue{}
	q.Init(createPeople(numPeople), &sync.Mutex{})
	task.End()
	return q
}

func InitializeCounters(ctx context.Context, num int) []*domain.Counter {
	ctx, task := trace.NewTask(ctx, "make counters")
	counters := make([]*domain.Counter, 0)
	for i := 0; i < num; i++ {
		counters = append(counters, &domain.Counter{Id: i})
	}
	task.End()
	return counters
}
