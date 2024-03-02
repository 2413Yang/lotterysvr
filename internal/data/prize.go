package data

import "github.com/2413Yang/lotterysvr/internal/biz"

type prizeRepo struct {
	data *Data
}

func NewPrizeRepo(data *Data) biz.PrizeRepo {
	return &prizeRepo{
		data: data,
	}
}
