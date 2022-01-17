package service

import (
	"context"
	"github.com/timfame/todo-list/internal/models"
)

type IService interface {
	GetLists(ctx context.Context) ([]models.List, error)
	AddList(ctx context.Context, list models.List) error
	RemoveList(ctx context.Context, listIndex int) error
	AddTask(ctx context.Context, listIndex int, task models.Task) error
	MarkTaskDone(ctx context.Context, listIndex, taskIndex int) error
}
