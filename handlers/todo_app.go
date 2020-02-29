package handlers

import (
	"context"
	gen "golang-todo-app/proto"

	"github.com/google/uuid"
)

// TodoAppServer exposes an API for interacting with todo items
type TodoAppServer struct {
	gen.UnimplementedTodoAppServer
}

// NewTodoAppServer returns a new todoAppServer
func NewTodoAppServer() *TodoAppServer {
	s := &TodoAppServer{}
	return s
}

// CreateTodo ...
func (h *TodoAppServer) CreateTodo(context.Context, *gen.CreateTodoRequest) (*gen.CreateTodoResponse, error) {
	return &gen.CreateTodoResponse{
		Title:       "Hello",
		Description: "Yo",
		Uuid:        uuid.New().String(),
	}, nil
}
