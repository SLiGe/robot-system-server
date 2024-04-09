//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"robot-system-server/internal/handler"
	"robot-system-server/internal/repository"
	"robot-system-server/internal/server"
	"robot-system-server/internal/service"
	"robot-system-server/pkg/app"
	"robot-system-server/pkg/jwt"
	"robot-system-server/pkg/log"
	"robot-system-server/pkg/server/http"
	"robot-system-server/pkg/sid"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewBeasenRepository,
	repository.NewFortuneRepository,
	repository.NewFortuneDataRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewBeasenService,
	service.NewFortuneService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewBeasenHandler,
	handler.NewFortuneHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {

	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
