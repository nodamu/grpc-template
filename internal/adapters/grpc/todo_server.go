package grpc

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	apiv1 "grpc-template/internal/api/v1"
	"grpc-template/internal/core/domain"
	"grpc-template/internal/core/ports"
)

type todoServer struct {
	todoService ports.TodoService
	apiv1.UnimplementedTodoServiceServer
}

func (s todoServer) CreateTodo(ctx context.Context, req *apiv1.CreateTodoRequest) (*apiv1.CreateTodoResponse, error) {
	todo := &domain.Todo{
		Title:       req.Title,
		Description: req.Description,
		IsCompleted: req.IsCompleted,
		CreatedAt:   req.CreatedAt.AsTime(),
	}
	id, err := s.todoService.CreateTodo(ctx, todo)

	if err != nil {
		log.Err(err).Msg("error occurred creating todo")
		return nil, status.Error(codes.Internal, "error occurred creating todo")
	}

	return &apiv1.CreateTodoResponse{Id: id}, nil

}
