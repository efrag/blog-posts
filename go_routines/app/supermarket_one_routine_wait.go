package app

import (
	"context"
	"runtime/trace"
	"strconv"
	"sync"

	"github.com/efrag/blog-posts/go_routines/domain"
)

func processQueueWait(ctx context.Context, q *domain.Queue, cs []*domain.Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	c := cs[0]

	for q.NumberOfPeople() > 0 {
		p := q.Pop()
		ctx, task := trace.NewTask(ctx, "process")
		trace.Log(ctx, "items", strconv.Itoa(p.Items))
		c.Process(p)
		task.End()
	}
}

func RunSupermarketGoRoutineWait(ctx context.Context, q *domain.Queue, cs []*domain.Counter) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go processQueueWait(ctx, q, cs, wg)
	wg.Wait()
}
