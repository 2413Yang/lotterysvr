package data

import (
	"github.com/2413Yang/lotterysvr/internal/conf"

	"github.com/2413Yang/pkg/midware/cache"
	"github.com/2413Yang/pkg/midware/gormcli"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDatabase, NewCache)

// Data .
type Data struct {
	// TODO wrapped database client
	db    *gorm.DB
	cache *cache.Client
}

// NewData .
func NewData(db *gorm.DB, cache *cache.Client) *Data {
	return &Data{db: db, cache: cache}
}

func NewDatabase(conf *conf.Data) *gorm.DB {
	dt := conf.GetDatabase()
	gormcli.Init(
		gormcli.WithAddr(dt.GetAddr()),
		gormcli.WithUser(dt.GetUser()),
		gormcli.WithPassword(dt.GetPassword()),
		gormcli.WithDataBase(dt.GetDatabase()),
		gormcli.WithMaxIdleConn(int(dt.GetMaxIdleConn())),
		gormcli.WithMaxOpenConn(int(dt.GetMaxOpenConn())),
		gormcli.WithMaxIdleTime(int64(dt.GetMaxIdleTime())),
		//如果设置了慢查询阈值，就打印日志
		gormcli.WithSlowThresholdMillisecond(dt.GetSlowThresholdMillisecond()),
	)
	return gormcli.GetDB()
}

func NewCache(conf *conf.Data) *cache.Client {
	dt := conf.GetRedis()
	cache.Init(
		cache.WithAddr(dt.GetAddr()),
		cache.WithPassWord(dt.GetPassword()),
		cache.WithDB(int(dt.GetDb())),
		cache.WithPoolSize(int(dt.GetPoolSize())),
	)
	return cache.GetRedisCli()
}
