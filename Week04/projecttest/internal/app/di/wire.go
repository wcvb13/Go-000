// +build wireinject

package di

import (
	"github.com/google/wire"
	"projecttest/internal/app/biz"
	"projecttest/internal/app/data"
	"projecttest/internal/app/service"
)

func InitApp() *App {
	panic(wire.Build(NewApp,service.NewUserService,biz.NewUserUsercase,data.Dataset))
}