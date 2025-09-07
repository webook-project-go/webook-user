package main

import (
	"context"
	v1 "github.com/webook-project-go/webook-apis/gen/go/apis/user/v1"
	_ "github.com/webook-project-go/webook-user/config"
	"github.com/webook-project-go/webook-user/ioc"
)

func main() {
	app := InitApp()
	shutdown := ioc.InitOTEL()
	defer shutdown(context.Background())
	v1.RegisterUserAuthServiceServer(app.Server, app.Service)
	err := app.Server.Serve()
	if err != nil {
		panic(err)
	}
}
