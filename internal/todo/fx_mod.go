package todo

import (
	"github.com/GeorgiyGusev/hack-backend/internal/todo/delivery"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"todo_module",
	fx.Invoke(func(e *echo.Echo) {
		delivery.Register(e)
	}),
)
