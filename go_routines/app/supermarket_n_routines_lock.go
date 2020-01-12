package app

import (
	"context"
	"runtime/trace"
	"strconv"
	"sync"
	"time"

	"github.com/efrag/blog-posts/go_routines/logger"

	"github.com/efrag/blog-posts/go_routines/domain"
)

func processQueueGoRLock(ctx context.Context, q *domain.Queue, c *domain.Counter, wg *sync.WaitGroup) {
	defer wg.Done()

	q.Lock.Lock()
	np := q.NumberOfPeople()
	time.Sleep(1 * time.Millisecond)
	p := q.Pop()
	q.Lock.Unlock()

	for np > 0 {
		ctx, task := trace.NewTask(ctx, "process")
		trace.Log(ctx, "items", strconv.Itoa(p.Items))
		trace.Log(ctx, "gorid", strconv.Itoa(int(logger.GetRoutineID())))
		c.Process(p)
		task.End()

		q.Lock.Lock()
		np = q.NumberOfPeople()
		time.Sleep(10 * time.Millisecond)
		p = q.Pop()
		q.Lock.Unlock()
	}
}

func RunSuperMarketGoRLock(ctx context.Context, q *domain.Queue, cs []*domain.Counter) {
	wg := &sync.WaitGroup{}

	for _, c := range cs {
		wg.Add(1)
		go processQueueGoRLock(ctx, q, c, wg)
	}
	wg.Wait()
}
