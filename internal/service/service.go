package service

import (
	pb "github.com/2413Yang/lotterysvr/api/lottery/v1"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewLotteryService)

type LotteryService struct {
	pb.UnimplementedLotteryServer
	// lotteryCase *biz.LotteryCase
	// limitCase   *biz.LimitCase
	// adminCase   *biz.AdminCase
}

func NewLotteryService() *LotteryService {
	return &LotteryService{}
}
