//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/webook-project-go/webook-user/grpc"
	"github.com/webook-project-go/webook-user/ioc"
	"github.com/webook-project-go/webook-user/repository"
	"github.com/webook-project-go/webook-user/repository/cache"
	"github.com/webook-project-go/webook-user/repository/dao"
	"github.com/webook-project-go/webook-user/service"
)

var thirdPartyProvider = wire.NewSet(
	ioc.InitDatabase,
	ioc.InitRedis,
	ioc.InitEtcd,
	ioc.InitGrpcServer,
)

var serviceSet = wire.NewSet(
	service.NewAuthBindingService,
	service.NewUserService,
	repository.New,
	repository.NewRepository,
	dao.NewDao,
	cache.NewUserCache,
)

func InitApp() *App {
	wire.Build(
		wire.Struct(new(App), "*"),
		grpc.NewService,
		thirdPartyProvider,
		serviceSet,
	)
	return nil
}
