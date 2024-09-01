package chapter5

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type LocalCache struct {
	sync.RWMutex
	storage map[string]Item
}

type Item struct {
	value  interface{}
	expire time.Time
}

func NewCache() *LocalCache {

	c := &LocalCache{
		storage: make(map[string]Item),
	}

	go func(lc *LocalCache) {
		for {
			//每秒检查一次
			time.Sleep(time.Second)
			for key, item := range c.storage {
				if item.expire.Before(time.Now()) {
					delete(c.storage, key)
					fmt.Println("delete key:", key)
				}
			}
		}
	}(c)

	return c
}

func (lc *LocalCache) Set(key string, value interface{}, expire int64) {
	lc.Lock()
	defer lc.Unlock()
	lc.storage[key] = Item{
		value:  value,
		expire: time.Now().Add(time.Duration(expire) * time.Second),
	}
}

func (lc *LocalCache) Get(key string) (interface{}, bool) {
	lc.RLock()
	defer lc.RUnlock()
	if val, ok := lc.storage[key]; ok {
		fmt.Println(key, val.value)
		if val.expire.Before(time.Now()) {
			return nil, false
		}
		return val.value, true
	}
	return nil, false
}

func (lc *LocalCache) Delete(key string) {
	lc.Lock()
	defer lc.Unlock()
	delete(lc.storage, key)
}

func TestCache(t *testing.T) {
	cache := NewCache()
	cache.Set("my_name", "gary", 5)

	go func() {
		for {
			t.Log(cache.Get("my_name"))
			time.Sleep(time.Second)
		}

	}()

	select {}
}

func TestCache2(t *testing.T) {
	cache := NewCache()
	cache.Set("my_name", "gary", 10)

	go func() {
		for {
			t.Log(cache.Get("my_name"))
			time.Sleep(time.Second)
		}

	}()

	ticker := time.NewTicker(time.Second * 3)
	select {
	case <-ticker.C:
		cache.Delete("my_name")
	}

	select {}
}
