package handlers

import (
	"context"
	gen "golang-todo-app/proto"
	"golang-todo-app/services"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
)

// TodoAppServer exposes an API for interacting with todo items
type TodoAppServer struct {
	gen.UnimplementedTodoAppServer
	todoService services.TodoService
}

// NewTodoAppServer returns a new todoAppServer
func NewTodoAppServer() *TodoAppServer {
	s := &TodoAppServer{
		todoService: services.NewTodoService(),
	}
	return s
}

// CreateTodo ...
func (h *TodoAppServer) CreateTodo(context context.Context, request *gen.CreateTodoRequest) (*gen.CreateTodoResponse, error) {
	grpc_zap.Extract(context).Info("Log log")

	return h.todoService.CreateTodo(context, request)
}
