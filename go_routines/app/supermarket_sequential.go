package app

import (
	"context"
	"runtime/trace"
	"strconv"

	"github.com/efrag/blog-posts/go_routines/domain"
)

func RunSupermarketSequential(ctx context.Context, q *domain.Queue, cs []*domain.Counter) {
	c := cs[0]

	for q.NumberOfPeople() > 0 {
		p := q.Pop()

		if p == nil {
			return
		}

		ctx, task := trace.NewTask(ctx, "process")
		trace.Log(ctx, "items", strconv.Itoa(p.Items))
		c.Process(p)
		task.End()
	}
}
