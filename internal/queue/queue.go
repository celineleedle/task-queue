package queue

import (
	"fmt"
	"sync"
	"time"

	"github.com/celineleedle/task-queue/internal/model"
)

type TaskQueue struct {
	tasks map[string]*model.Task

	// priority -> queue of task ids
	queues map[model.Priority][]string

	// concurrency
	lock *sync.RWMutex
	cond *sync.Cond

	closed bool
}

func NewTaskQueue() *TaskQueue {
	lock := &sync.RWMutex{}

	highPriority := []string{}
	medPriority := []string{}
	lowPriority := []string{}

	queues := map[model.Priority][]string{
		model.PriorityHigh: highPriority,
		model.PriorityMed:  medPriority,
		model.PriorityLow:  lowPriority,
	}

	return &TaskQueue{
		tasks:  map[string]*model.Task{},
		queues: queues,
		lock:   lock,
		cond:   sync.NewCond(lock),
		closed: false,
	}
}

func (t *TaskQueue) Enqueue(task *model.Task) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	if t.closed {
		return fmt.Errorf("can not enqueue when task queue is closed")
	}

	task.Status = model.StatusPending
	task.CreatedAt = time.Now()

	t.tasks[task.ID] = task

	_, ok := t.queues[task.Priority]
	if !ok {
		return fmt.Errorf("invalid priority value for task: %q", task.Priority.String())
	}
	t.queues[task.Priority] = append(t.queues[task.Priority], task.ID)

	t.cond.Signal()

	return nil
}

func (t *TaskQueue) Dequeue() (*model.Task, error) {
	t.lock.Lock()
	defer t.lock.Unlock()

	for {
		if t.closed {
			return nil, fmt.Errorf("can not dequeue when task queue is closed")
		}

		id := ""
		for _, p := range []model.Priority{model.PriorityHigh, model.PriorityMed, model.PriorityLow} {
			if len(t.queues[p]) > 0 {
				id = t.queues[p][0]
				t.queues[p] = t.queues[p][1:]
				break
			}
		}
		if id == "" {
			t.cond.Wait()
			continue
		}

		task := t.tasks[id]
		task.Status = model.StatusProcessing

		now := time.Now()
		task.StartedAt = &now

		task.Tries++

		return task, nil
	}
}

func (t *TaskQueue) Complete(id string, result string) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	task, ok := t.tasks[id]
	if !ok {
		return fmt.Errorf("task not found with id: %q", id)
	}

	task.Status = model.StatusCompleted

	now := time.Now()
	task.CompletedAt = &now

	task.Result = result

	return nil
}

func (t *TaskQueue) Fail(id string, err string) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	task, ok := t.tasks[id]
	if !ok {
		return fmt.Errorf("task not found with id: %q", id)
	}

	if task.Tries < task.MaxTries {
		task.Status = model.StatusPending
		task.Error = err
		task.StartedAt = nil

		t.queues[task.Priority] = append(t.queues[task.Priority], task.ID)
		t.cond.Signal()
	} else {
		task.Status = model.StatusFailed
		task.Error = err

		now := time.Now()
		task.CompletedAt = &now
	}

	return nil
}

func (t *TaskQueue) Get(id string) (*model.Task, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	task, ok := t.tasks[id]
	if !ok {
		return nil, fmt.Errorf("task not found with id: %q", id)
	}
	return task, nil
}

func (t *TaskQueue) Close() {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.closed = true
	t.cond.Broadcast()
}

func (t *TaskQueue) List() []*model.Task {
	t.lock.RLock()
	defer t.lock.RUnlock()

	tasks := []*model.Task{}

	for _, task := range t.tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

func (t *TaskQueue) Stats() Stats {
	t.lock.RLock()
	defer t.lock.RUnlock()

	numPending, numProcessing, numCompleted, numFailed := 0, 0, 0, 0
	for _, task := range t.tasks {
		switch task.Status {
		case model.StatusPending:
			numPending++
		case model.StatusProcessing:
			numProcessing++
		case model.StatusCompleted:
			numCompleted++
		case model.StatusFailed:
			numFailed++
		}
	}

	return Stats{
		NumTasks:      len(t.tasks),
		NumPending:    numPending,
		NumProcessing: numProcessing,
		NumCompleted:  numCompleted,
		NumFailed:     numFailed,
	}
}
