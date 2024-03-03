//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/2413Yang/lotterysvr/internal/biz"
	"github.com/2413Yang/lotterysvr/internal/conf"
	"github.com/2413Yang/lotterysvr/internal/data"
	"github.com/2413Yang/lotterysvr/internal/interfaces"
	"github.com/2413Yang/lotterysvr/internal/server"
	"github.com/2413Yang/lotterysvr/internal/service"
	"github.com/2413Yang/lotterysvr/internal/task"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		interfaces.ProviderSet,
		task.ProviderSet,
		newApp))
}
