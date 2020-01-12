package app

import (
	"context"
	"runtime/trace"
	"strconv"
	"sync"
	"time"

	"github.com/efrag/blog-posts/go_routines/domain"
)

func processQueueRoutines(ctx context.Context, q *domain.Queue, c *domain.Counter, wg *sync.WaitGroup) {
	defer wg.Done()

	for q.NumberOfPeople() > 0 {
		time.Sleep(10 * time.Millisecond)

		p := q.Pop()
		if p != nil {
			ctx, task := trace.NewTask(ctx, "process")
			trace.Log(ctx, "items", strconv.Itoa(p.Items))
			c.Process(p)
			task.End()
		}
	}
}

func RunSupermarketNGoRoutines(ctx context.Context, q *domain.Queue, cs []*domain.Counter) {
	wg := &sync.WaitGroup{}

	for _, c := range cs {
		wg.Add(1)
		go processQueueRoutines(ctx, q, c, wg)
	}
	wg.Wait()
}
