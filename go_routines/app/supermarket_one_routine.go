package app

import (
	"context"
	"runtime/trace"
	"strconv"

	"github.com/efrag/blog-posts/go_routines/domain"
)

func processQueue(ctx context.Context, q *domain.Queue, cs []*domain.Counter) {
	c := cs[0]

	for q.NumberOfPeople() > 0 {
		p := q.Pop()
		ctx, task := trace.NewTask(ctx, "process")
		trace.Log(ctx, "items", strconv.Itoa(p.Items))
		c.Process(p)
		task.End()
	}
}

func RunSupermarketGoRoutine(ctx context.Context, q *domain.Queue, cs []*domain.Counter) {
	// this will return immediately and the outcome of the code is
	// not going to be visible at all since the main function will
	// exit without waiting for the go routine to return
	go processQueue(ctx, q, cs)
}
