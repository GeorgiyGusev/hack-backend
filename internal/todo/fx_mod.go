package todo

import (
	"github.com/GeorgiyGusev/auth-library/provider"
	"github.com/GeorgiyGusev/hack-backend/internal/todo/delivery"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"todo_module",
	fx.Invoke(func(e *echo.Echo, provider provider.AuthProvider) {
		delivery.Register(e, provider)
	}),
)
