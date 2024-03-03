package task

import (
	"context"
	"fmt"
	"time"

	"github.com/2413Yang/lotterysvr/internal/constant"
	"github.com/2413Yang/lotterysvr/internal/utils"
	"github.com/2413Yang/pkg/midware/cache"
	"github.com/go-kratos/kratos/v2/log"
)

func DoResetUserLotteryNumsTask() {
	go ResetUserLotteryNums()
}

func ResetUserLotteryNums() {
	//log.Infof("重置今日用户抽奖次数")
	for i := 0; i < constant.IpFrameSize; i++ {
		key := fmt.Sprintf(constant.UserLotteryDayNumPrefix+"%d", i)
		if err := cache.GetRedisCli().Delete(context.Background(), key); err != nil {
			log.Errorf("ResetIPLotteryNums err:%v", err)
		}
	}

	// IP当天的统计数，整点归零，设置定时器
	duration := utils.NextDayDuration()
	time.AfterFunc(duration, ResetUserLotteryNums) //等待时间段d过去，然后调用func
}
