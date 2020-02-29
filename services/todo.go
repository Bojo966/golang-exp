package services

import (
	"context"
	gen "golang-todo-app/proto"

	"github.com/google/uuid"
)

// TodoService ...
type TodoService interface {
	CreateTodo(context.Context, *gen.CreateTodoRequest) (*gen.CreateTodoResponse, error)
}

type todoService struct {
}

// NewTodoService ...
func NewTodoService() TodoService {
	return todoService{}
}

func (s todoService) CreateTodo(context.Context, *gen.CreateTodoRequest) (*gen.CreateTodoResponse, error) {
	return &gen.CreateTodoResponse{
		Title:       "Hello",
		Description: "Yo",
		Uuid:        uuid.New().String(),
	}, nil
}
