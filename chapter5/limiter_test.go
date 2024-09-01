package chapter5

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"testing"
	"time"
)

type Limiter struct {
	Key    string
	Rate   int //频率
	Period int //秒级限制器
	rds    *redis.Client
}

func NewLimiter(rds *redis.Client, key string, rate int, period int) *Limiter {
	limiter := Limiter{
		Key:    key,
		Rate:   rate,
		Period: period,
		rds:    rds,
	}
	return &limiter
}

func (l *Limiter) Allow() bool {
	ctx := context.Background()
	timestamp := time.Now().Unix()
	suffix := timestamp / int64(l.Period)
	key := fmt.Sprintf("%s:%d", l.Key, suffix) //name:1771234567
	count, err := l.rds.Incr(ctx, key).Result()
	if err != nil {
		return false
	}
	if count == 1 {
		err = l.rds.Expire(ctx, key, time.Duration(l.Period+1)*time.Second).Err()
		if err != nil {
			return false
		}
	}
	if count > int64(l.Rate) {
		return false
	}
	return true
}

func TestAllow(t *testing.T) {
	rds := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:16379",
		Password: "123456",
		DB:       0,
	})
	limiter := NewLimiter(rds, "test", 10, 10)
	i := 1
	for {
		if limiter.Allow() {
			log.Println(i, "allow")
			//fmt.Println("allow")
		} else {
			log.Println(i, "deny")
			//fmt.Println("deny")
		}
		i++
		time.Sleep(time.Millisecond * 500)
	}
}
