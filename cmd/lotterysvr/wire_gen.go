// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data) (*kratos.App, func(), error) {
	db := data.NewDatabase(confData)
	client := data.NewCache(confData)
	dataData := data.NewData(db, client)
	prizeRepo := data.NewPrizeRepo(dataData)
	couponRepo := data.NewCouponRepo(dataData)
	blackUserRepo := data.NewBlackUserRepo(dataData)
	blackIpRepo := data.NewBlackIpRepo(dataData)
	resultRepo := data.NewResultRepo(dataData)
	transaction := data.NewTransaction(dataData)
	lotteryCase := biz.NewLotteryCase(prizeRepo, couponRepo, blackUserRepo, blackIpRepo, resultRepo, transaction)
	lotteryTimesRepo := data.NewLotteryTimesRepo(dataData)
	limitCase := biz.NewLimitCase(blackUserRepo, blackIpRepo, lotteryTimesRepo, transaction)
	adminCase := biz.NewAdminCase(prizeRepo, couponRepo, lotteryTimesRepo, resultRepo)
	lotteryService := service.NewLotteryService(lotteryCase, limitCase, adminCase)
	grpcServer := server.NewGRPCServer(confServer, lotteryService)
	adminService := service.NewAdminService(adminCase)
	handler := interfaces.NewHandler(lotteryService, adminService)
	httpServer := server.NewHTTPServer(confServer, handler)
	taskServer := task.NewTaskServer(lotteryService, confServer)
	app := newApp(grpcServer, httpServer, taskServer)
	return app, func() {
	}, nil
}
