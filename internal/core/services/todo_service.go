package services

import (
	"context"
	"grpc-template/internal/core/domain"
	"grpc-template/internal/core/ports"
)

type TodoService struct {
	todoRepository ports.TodoRepository
}

func (t TodoService) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	err := t.todoRepository.CreateTodo(ctx, todo)
	if err != nil {
		return err
	}
	return nil
}
