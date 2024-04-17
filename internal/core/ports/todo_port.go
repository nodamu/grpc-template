package ports

import (
	"context"
	"grpc-template/internal/core/domain"
)

type TodoService interface {
	CreateTodo(ctx context.Context, todo *domain.Todo) (string, error)
}

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *domain.Todo) (string, error)
}
