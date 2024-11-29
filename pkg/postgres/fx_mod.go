package postgres

import "go.uber.org/fx"

var Module = fx.Module(
	"postgres",
	fx.Provide(LoadConfig, NewPostgressConn),
	fx.Invoke(TestPostgreConn),
)
