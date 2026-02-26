package worker

import (
	"context"
	"fmt"
	"sync"

	"github.com/celineleedle/task-queue/internal/queue"
)

type WorkerPool struct {
	numWorkers int
	queue      *queue.TaskQueue
	handlers   map[string]Handler
	wg         *sync.WaitGroup
}

func NewWorkerPool(numWorkers int, queue *queue.TaskQueue) *WorkerPool {
	wg := &sync.WaitGroup{}

	return &WorkerPool{
		numWorkers: numWorkers,
		queue:      queue,
		handlers:   map[string]Handler{},
		wg:         wg,
	}
}

func (wp *WorkerPool) RegisterHandler(taskType string, handler Handler) {
	wp.handlers[taskType] = handler
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

func (wp *WorkerPool) Start(ctx context.Context) {
	for range wp.numWorkers {
		wp.wg.Add(1)

		go func() {
			defer wp.wg.Done()

			for {
				// check context for cancellation
				select {
				case <-ctx.Done():
					return
				default:
				}

				task, err := wp.queue.Dequeue()
				if err != nil {
					return
				}

				handler, ok := wp.handlers[task.Type]
				if !ok {
					wp.queue.Fail(task.ID, fmt.Sprintf("No handler registered for task type: %q", task.Type))
					continue
				}

				res, err := handler(ctx, task)
				if err != nil {
					wp.queue.Fail(task.ID, err.Error())
					continue
				}

				wp.queue.Complete(task.ID, res)
			}

		}()
	}
}
