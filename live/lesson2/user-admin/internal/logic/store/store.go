package store

import (
	"context"
	"github.com/gogf/gf/v2/os/gcache"
	"user-admin/internal/model"
	"user-admin/internal/service"
)

type (
	sStore struct {
		Cache *gcache.Cache
	}
)

var ctx = context.Background()

func init() {
	service.RegisterStore(New())
}

func New() *sStore {
	return &sStore{
		Cache: gcache.New(),
	}
}

func (s *sStore) GetUser(userName string) (user model.User, err error) {
	res, err := s.Cache.Get(ctx, userName)
	err = res.Struct(&user)
	return
}

func (s *sStore) SetUser(user model.User) (err error) {
	return s.Cache.Set(ctx, user.UserName, user, 0)
}
