package service

import (
	"context"
	"fmt"

	pb "github.com/2413Yang/lotterysvr/api/lottery/v1"
	"github.com/2413Yang/lotterysvr/internal/constant"
	"github.com/go-kratos/kratos/v2/log"
)

func (s *LotteryService) LotteryV1(ctx context.Context, req *pb.LotteryReq) (*pb.LotteryRsp, error) {

	rsp := &pb.LotteryRsp{
		CommonRsp: &pb.CommonRspInfo{
			Code:   int32(Success),
			Msg:    GetErrMsg(Success),
			UserId: req.UserId,
		},
	}
	defer func() {
		//通过对应的Code，获取msg
		rsp.CommonRsp.Msg = GetErrMsg(ErrCode(rsp.CommonRsp.Code))
	}()

	userID := uint(req.UserId)
	log.Infof("LotteryV1|user_id=%d", userID)
	_ = fmt.Sprintf(constant.LotterLockKeyPrefix+"%d", userID)

	return nil, nil
}
