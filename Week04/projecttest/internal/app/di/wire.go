package di

import (
	"github.com/google/wire"
	"projecttest/internal/app/biz"
	"projecttest/internal/app/data"
	"projecttest/internal/app/service"
)

func InitApp() *App {
	panic(wire.Build(data.Dataset,biz.NewUserUsercase,service.NewUserService))
}