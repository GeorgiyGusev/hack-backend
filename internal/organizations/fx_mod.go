package organizations

import (
	"github.com/GeorgiyGusev/hack-backend/internal/organizations/delivery"
	"github.com/GeorgiyGusev/hack-backend/internal/organizations/storage/repository"
	"github.com/GeorgiyGusev/hack-backend/internal/organizations/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"organizations",
	fx.Provide(
		fx.Annotate(usecase.NewImpl, fx.As(new(usecase.Usecase))),
		fx.Annotate(repository.NewRepository, fx.As(new(usecase.Repository))),
	),
	fx.Invoke(delivery.Register),
)
