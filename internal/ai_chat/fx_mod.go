package ai_chat

import (
	"github.com/GeorgiyGusev/hack-backend/internal/ai_chat/delivery"
	"github.com/GeorgiyGusev/hack-backend/internal/ai_chat/gateway"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"ai_chat",
	fx.Provide(
		fx.Annotate(gateway.NewGateway),
	),
	fx.Invoke(delivery.Register),
)
