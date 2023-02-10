package lib

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewEchoHandler),
	fx.Provide(NewDatabase),
)

type RepositoryImpl struct {
	Database
}
