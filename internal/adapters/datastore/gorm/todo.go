package gorm

import (
	"context"
	"grpc-template/internal/core/domain"

	"gorm.io/gorm"
)

type TodoPGRepo struct {
	db *gorm.DB
}

// CreateTodo is a concrete implementation of the CreateTodo method in the Todo Repository interface
func (s TodoPGRepo) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	err := s.db.Create(todo).Error
	if err != nil {
		return err
	}
	return nil
}
