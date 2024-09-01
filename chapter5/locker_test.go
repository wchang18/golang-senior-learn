package chapter5

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"testing"
	"time"
)

type Locker struct {
	rds *redis.Client
	key string
	ctx context.Context
}

func NewLock(rds *redis.Client, key string) *Locker {
	return &Locker{
		rds: rds,
		key: key,
		ctx: context.Background(),
	}
}

func (r *Locker) Lock() {
	for {
		ok := r.rds.SetNX(r.ctx, r.key, 1, time.Minute*5)

		if ok.Val() {
			break
		} else {
			time.Sleep(time.Second)
		}
	}
}

func (r *Locker) UnLock() {
	r.rds.Del(r.ctx, r.key)
}

func TestLocker(t *testing.T) {
	rds := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:16379",
		Password: "123456",
		DB:       1,
	})

	key := "order:123"
	var wg sync.WaitGroup
	fmt.Println(rds, key)

	go func() {
		wg.Add(1)
		defer wg.Done()
		logx.Info("回调开始")
		//locker := NewLock(rds, key)
		//locker.Lock()
		//defer locker.UnLock()
		time.Sleep(time.Second * 3)
		fmt.Println("回调订单执行完成")
		logx.Info("回调结束")
	}()

	time.Sleep(time.Millisecond * 100)

	go func() {
		wg.Add(1)
		defer wg.Done()
		logx.Info("脚本开始")
		//locker := NewLock(rds, key)
		//locker.Lock()
		//defer locker.UnLock()
		time.Sleep(time.Second)
		fmt.Println("脚本订单执行完成")
		logx.Info("脚本结束")
	}()

	wg.Wait()
}
