package chapter5

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
	"testing"
	"time"
)

var redisClient *redis.Client

func GetRedisClient() *redis.Client {
	var once sync.Once
	once.Do(func() {
		rds := redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:16379",
			Password: "123456",
		})
		redisClient = rds
	})
	return redisClient
}

func TestString(t *testing.T) {
	rds := GetRedisClient()
	ctx := context.Background()
	status := rds.Set(ctx, "name1", "tom1", 0)

	status = rds.Set(ctx, "name2", "tom2", time.Second*20)
	t.Log(status.Val())

	is := rds.SetNX(ctx, "name3", "tom3", 0)
	t.Log(is.Val())
	//
	str := rds.Get(ctx, "name3")
	t.Log(str.Val())
	//
	rds.Expire(ctx, "name3", time.Second*10)
	//
	j := rds.Del(ctx, "name1")
	t.Log(j.Val())
	//
	i := rds.Exists(ctx, "name4")
	t.Log(i.Val())
	//
	rds.Set(ctx, "number1", 1, 0)
	rds.Incr(ctx, "number1")
	rds.Incr(ctx, "number1")
	t.Log(rds.Get(ctx, "number1").Val())
	rds.Decr(ctx, "number1")
	t.Log(rds.Get(ctx, "number1").Val())
}

func TestHash(t *testing.T) {
	rds := GetRedisClient()
	ctx := context.Background()
	var m1 map[string]string
	m1 = map[string]string{
		"name1": "tom1",
		"name2": "tom2",
		"name3": "tom3",
	}
	t.Logf("%+v", m1)

	rds.HSet(ctx, "user1", "name1", "tom1")
	rds.HSet(ctx, "user1", "name2", "tom2")
	rds.HSet(ctx, "user1", "name3", "tom3")

	t.Log(rds.HGet(ctx, "user1", "name1").Val())

	rds.HMSet(ctx, "user2", "name1", "jim1", "name2", "jim2", "name3", "jim3")
	t.Log(rds.HGetAll(ctx, "user2").Val())

	t.Log(rds.HDel(ctx, "user1", "name1").Val())
	t.Log(rds.HGetAll(ctx, "user1").Val())
	t.Log(rds.HExists(ctx, "user1", "name1").Val())
}

func TestList(t *testing.T) {
	rds := GetRedisClient()
	ctx := context.Background()

	var list []string
	list = []string{"a", "b", "c"}
	fmt.Println(list)

	rds.LPush(ctx, "list1", "a", "b", "c")

	rds.Expire(ctx, "list1", time.Second*5)

	t.Log(rds.LRange(ctx, "list1", 0, 10).Val())

	t.Log(rds.LLen(ctx, "list1").Val())

	t.Log(rds.LPop(ctx, "list1").Val())

	t.Log(rds.LRange(ctx, "list1", 0, 10).Val())
	//
	rds.RPush(ctx, "list2", "e", "f", "g")
	rds.Expire(ctx, "list2", time.Second*5)
	t.Log(rds.LRange(ctx, "list2", 0, 10).Val())
	t.Log(rds.RPop(ctx, "list2").Val())
	t.Log(rds.LRange(ctx, "list2", 0, 10).Val())
}

func TestSet(t *testing.T) {
	rds := GetRedisClient()
	ctx := context.Background()
	rds.SAdd(ctx, "set1", "a", "b", "c", "c")
	t.Log(rds.SMembers(ctx, "set1").Val())

	rds.SAdd(ctx, "set2", "d", "e", "f", "c")
	t.Log(rds.SMembers(ctx, "set2").Val())

	t.Log(rds.SIsMember(ctx, "set1", "c").Val())
	t.Log(rds.SIsMember(ctx, "set2", "g").Val())
	//
	t.Log(rds.SInter(ctx, "set1", "set2").Val())

	t.Log(rds.SUnion(ctx, "set1", "set2").Val())
}

func TestSortedSet(t *testing.T) {
	rds := GetRedisClient()
	ctx := context.Background()
	rds.ZAdd(ctx, "zset1", redis.Z{
		Score:  1,
		Member: "a",
	}, redis.Z{
		Score:  2,
		Member: "b",
	}, redis.Z{
		Score:  30,
		Member: "c",
	})
	t.Log(rds.ZRange(ctx, "zset1", 0, 10).Val())

	rds.ZAdd(ctx, "zset1", redis.Z{
		Score:  3,
		Member: "d",
	})
	t.Log(rds.ZRange(ctx, "zset1", 0, 10).Val())
}
