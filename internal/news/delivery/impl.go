package delivery

import (
	"github.com/labstack/echo/v4"
	"log/slog"
)

//go:generate oapi-codegen --config=../gen/gen.cfg.yaml https://raw.githubusercontent.com/GeorgiyGusev/hack-api/refs/heads/main/service-api.yaml
type Handlers struct {
	logger *slog.Logger
	//usecase usecase.Usecase
}

func (h Handlers) GetAllNews(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h Handlers) CreateNews(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h Handlers) GetNewsById(ctx echo.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
