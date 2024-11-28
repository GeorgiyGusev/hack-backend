package delivery

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

//go:generate oapi-codegen --config=./gen.cfg.yaml ../../../api/api.yaml
type Handlers struct {
	Todos map[string]Todo
}

func Register(e EchoRouter) {
	impl := &Handlers{
		Todos: make(map[string]Todo),
	}

	RegisterHandlers(e, impl)
}

func (h *Handlers) GetTodos(ctx echo.Context, _ GetTodosParams) error {
	var resp []Todo
	for _, v := range h.Todos {
		resp = append(resp, v)
	}
	if resp == nil {
		return ctx.JSON(200, []Todo{})
	}
	return ctx.JSON(200, resp)
}

func (h *Handlers) CreateTodo(ctx echo.Context) error {
	var req *CreateTodo
	if err := ctx.Bind(&req); err != nil {
		log.Println(err)
	}

	id := uuid.NewString()
	defaultStatus := TodoStatusPending
	createdAt := time.Now()
	updatedAt := createdAt
	h.Todos[id] = Todo{
		Id:          &id,
		Title:       &req.Title,
		Description: req.Description,
		Status:      &defaultStatus,
		CreatedAt:   &createdAt,
		UpdatedAt:   &updatedAt,
	}

	return ctx.JSON(200, h.Todos[id])
}

func (h *Handlers) DeleteTodo(ctx echo.Context, id string) error {
	delete(h.Todos, id)
	return ctx.NoContent(204)
}

func (h *Handlers) GetTodoById(ctx echo.Context, id string) error {
	return ctx.JSON(200, h.Todos[id])
}

func (h *Handlers) UpdateTodo(ctx echo.Context, id string) error {
	var req *UpdateTodo
	if err := ctx.Bind(&req); err != nil {
		log.Println(err)
		return ctx.JSON(500, Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	updatedAt := time.Now()
	todo := h.Todos[id]
	todo.Title = req.Title
	todo.Description = req.Description
	todo.Status = (*TodoStatus)(req.Status)
	todo.UpdatedAt = &updatedAt
	h.Todos[id] = todo

	return ctx.JSON(200, h.Todos[id])
}
