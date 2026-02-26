package worker

import (
	"context"

	"github.com/celineleedle/task-queue/internal/model"
)

type Handler func(ctx context.Context, task *model.Task) (string, error)
