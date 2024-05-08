package gorm

import (
	"context"
	"grpc-template/internal/core/domain"

	"gorm.io/gorm"
)

type TodoPGRepo struct {
	db *gorm.DB
}

// SaveTodo is a concrete implementation of the SaveTodo method in the Todo Repository interface
func (s TodoPGRepo) SaveTodo(ctx context.Context, todo *domain.Todo) (string, error) {
	err := s.db.Create(todo).Error
	if err != nil {
		return "", nil
	}
	return todo.ID, nil
}
