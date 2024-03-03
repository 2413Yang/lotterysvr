package service

import (
	pb "github.com/2413Yang/lotterysvr/api/lottery/v1"
	"github.com/2413Yang/lotterysvr/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewLotteryService, NewAdminService)

type LotteryService struct {
	pb.UnimplementedLotteryServer
	lotteryCase *biz.LotteryCase
	limitCase   *biz.LimitCase
	adminCase   *biz.AdminCase
}

func NewLotteryService(loc *biz.LotteryCase, lc *biz.LimitCase, ac *biz.AdminCase) *LotteryService {
	return &LotteryService{
		lotteryCase: loc,
		limitCase:   lc,
		adminCase:   ac,
	}
}

// AdminService 奖品管理后台
type AdminService struct {
	adminCase *biz.AdminCase
}

func NewAdminService(ac *biz.AdminCase) *AdminService {
	return &AdminService{
		adminCase: ac,
	}
}
