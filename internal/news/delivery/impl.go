package delivery

import (
	"context"
	"github.com/GeorgiyGusev/hack-backend/internal/news/gen"
	"github.com/GeorgiyGusev/hack-backend/internal/news/usecase"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

//go:generate oapi-codegen --config=../gen/gen.cfg.yaml https://raw.githubusercontent.com/GeorgiyGusev/hack-api/refs/heads/main/service-api.yaml
type Handlers struct {
	logger  *slog.Logger
	usecase usecase.Usecase
}

func (h *Handlers) GetAllNews(ctx echo.Context) error {
	news, err := h.usecase.GetAllNews(context.Background())
	if err != nil {
		h.logger.Error("Cannot get all news", "error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	resp := []gen.News{}
	for _, v := range *news {
		resp = append(resp, gen.News{
			Id:             v.Id,
			Title:          v.Title,
			Description:    v.Description,
			PublishData:    v.CreatedAt.String(),
			OrganizationId: v.OrganizationId,
			Media:          v.Media,
		})
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (h *Handlers) CreateNews(ctx echo.Context) error {
	var req gen.CreateNews
	if err := ctx.Bind(&req); err != nil {
		h.logger.Error("Cannot bind request", "error", err.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	err := h.usecase.Create(context.Background(), &req)
	if err != nil {
		h.logger.Error("Cannot create news", "error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusCreated)
}

func (h *Handlers) GetNewsById(ctx echo.Context, id string) error {
	news, err := h.usecase.GetByID(context.Background(), id)
	if err != nil {
		h.logger.Error("Cannot get news", "error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, gen.News{
		Id:             news.Id,
		Title:          news.Title,
		Description:    news.Description,
		OrganizationId: news.OrganizationId,
		Media:          news.Media,
		PublishData:    news.CreatedAt.String(),
	})
}
