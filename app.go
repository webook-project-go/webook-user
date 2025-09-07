package main

import (
	"github.com/webook-project-go/webook-pkgs/grpcx"
	"github.com/webook-project-go/webook-user/grpc"
)

type App struct {
	Service *grpc.Service
	Server  *grpcx.GrpcxServer
}
