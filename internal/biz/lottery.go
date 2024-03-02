package biz

type LotteryCase struct {
	prizeRepo     PrizeRepo
	couponRepo    CouponRepo
	blackUserRepo BlackUserRepo
	blackIpRepo   BlackIpRepo
	resultRepo    ResultRepo
	tm            Transaction
}

func NewLotteryCase(pr PrizeRepo, cr CouponRepo, bur BlackUserRepo,
	bir BlackIpRepo, result ResultRepo, tm Transaction) *LotteryCase {
	return &LotteryCase{
		prizeRepo:     pr,
		couponRepo:    cr,
		blackUserRepo: bur,
		blackIpRepo:   bir,
		resultRepo:    result,
		tm:            tm,
	}
}
