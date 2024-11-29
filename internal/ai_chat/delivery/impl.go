package delivery

import (
	"context"
	"github.com/GeorgiyGusev/hack-backend/internal/ai_chat/gateway"
	"github.com/GeorgiyGusev/hack-backend/internal/ai_chat/gen"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

//go:generate oapi-codegen --config=../gen/gen.cfg.yaml https://raw.githubusercontent.com/GeorgiyGusev/hack-api/refs/heads/main/service-api.yaml
type Handlers struct {
	logger  *slog.Logger
	gateway *gateway.Gateway
}

func Register(logger *slog.Logger, gateway *gateway.Gateway, ec *echo.Echo) {
	impl := Handlers{
		logger:  logger,
		gateway: gateway,
	}
	gen.RegisterHandlers(ec, &impl)
}

func (h *Handlers) AiChat(ctx echo.Context) error {
	var req gen.AiChatMessage
	if err := ctx.Bind(&req); err != nil {
		h.logger.Error("Cannot bind req", "error", err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	resp, err := h.gateway.SendMessage(context.Background(), req.Message)
	if err != nil {
		h.logger.Error("Cannot send message", "error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, gen.AiChatMessage{
		Message: resp,
	})
}
