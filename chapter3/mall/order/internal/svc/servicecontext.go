package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"golang-senior-learn/chapter3/mall/order/internal/config"
	"golang-senior-learn/chapter3/mall/pay/pay"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
)

type ServiceContext struct {
	Config config.Config
	PayRpc pay.Pay
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: pay.NewPay(zrpc.MustNewClient(c.PayRpc)),
		DB:     NewDb(c),
	}
}

var DB *gorm.DB

func NewDb(c config.Config) *gorm.DB {
	var once sync.Once
	once.Do(func() {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{},
		)

		db, err := gorm.Open(mysql.Open(c.MysqlDb), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			panic(err)
		}
		DB = db
	})
	return DB
}
