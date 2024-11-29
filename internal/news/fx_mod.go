package news

import (
	"github.com/GeorgiyGusev/hack-backend/internal/news/delivery"
	"github.com/GeorgiyGusev/hack-backend/internal/news/storage/repository"
	"github.com/GeorgiyGusev/hack-backend/internal/news/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"news",
	fx.Provide(
		fx.Annotate(usecase.NewImpl, fx.As(new(usecase.Usecase))),
		fx.Annotate(repository.NewRepository, fx.As(new(usecase.Repository))),
	),
	fx.Invoke(delivery.Register),
)
