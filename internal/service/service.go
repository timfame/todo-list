package service

import (
	"context"
	"fmt"
	"github.com/timfame/todo-list/internal/models"
	"sync"
)

type service struct {
	lists []models.List
	mu    sync.RWMutex
}

func New() *service {
	return &service{
		lists: make([]models.List, 0),
		mu:    sync.RWMutex{},
	}
}

func (s *service) GetLists(ctx context.Context) ([]models.List, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.lists, nil
}

func (s *service) AddList(ctx context.Context, list models.List) error {
	s.mu.Lock()
	list.ID = len(s.lists)+1
	s.lists = append(s.lists, list)
	s.mu.Unlock()
	return nil
}

func (s *service) RemoveList(ctx context.Context, listIndex int) error {
	s.mu.Lock()
	if listIndex < 0 || listIndex >= len(s.lists) {
		return fmt.Errorf("wrong list index")
	}
	for i := listIndex; i < len(s.lists)-1; i++ {
		s.lists[i] = s.lists[i+1]
	}
	s.lists = s.lists[:len(s.lists)-1]
	s.mu.Unlock()
	return nil
}

func (s *service) AddTask(ctx context.Context, listIndex int, task models.Task) error {
	s.mu.Lock()
	if listIndex < 0 || listIndex >= len(s.lists) {
		return fmt.Errorf("wrong list index")
	}
	s.lists[listIndex].Tasks = append(s.lists[listIndex].Tasks, task)
	s.mu.Unlock()
	return nil
}

func (s *service) MarkTaskDone(ctx context.Context, listIndex, taskIndex int) error {
	s.mu.Lock()
	if listIndex < 0 || listIndex >= len(s.lists) {
		return fmt.Errorf("wrong list index")
	}
	if taskIndex < 0 || taskIndex >= len(s.lists[listIndex].Tasks) {
		return fmt.Errorf("wrong task index")
	}
	s.lists[listIndex].Tasks[taskIndex].IsDone = !s.lists[listIndex].Tasks[taskIndex].IsDone
	s.mu.Unlock()
	return nil
}
