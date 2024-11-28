package main

import (
	authLib "github.com/GeorgiyGusev/auth-library"
	"github.com/GeorgiyGusev/hack-backend/internal/todo"
	httpSrvLib "github.com/GeorgiyGusev/http-srv-library"
	"github.com/go-playground/validator/v10"
	loggingLib "github.com/neiasit/logging-library"
	redisLib "github.com/neiasit/redis-library"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log/slog"
)

func main() {
	app := fx.New(

		// setting validator
		fx.Provide(func() *validator.Validate {
			return validator.New(
				validator.WithRequiredStructEnabled(),
			)
		}),

		// including platform libs here
		loggingLib.Module,
		authLib.AuthKeycloakModule,
		redisLib.Module,
		httpSrvLib.ModuleWithAuth,

		// setting logger
		fx.WithLogger(func(logger *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{
				Logger: logger,
			}
		}),

		// setup domains
		todo.Module,
	)

	app.Run()
}
