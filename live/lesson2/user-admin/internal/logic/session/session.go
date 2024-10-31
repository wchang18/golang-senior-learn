package session

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"user-admin/internal/service"
)

type (
	sSession struct {
		Session *ghttp.Session
	}
)

var session *ghttp.Session

func New() *sSession {
	return &sSession{
		Session: session,
	}
}

func init() {
	service.RegisterSession(New())
}

func (s *sSession) Init(req *ghttp.Request) {
	s.Session = req.Session
}

func (s *sSession) Get(key string) interface{} {
	res, _ := s.Session.Get(key)
	return res
}

func (s *sSession) Set(key string, value interface{}) {
	s.Session.Set(key, value)
}

func (s *sSession) Remove(key string) {
	s.Session.Remove(key)
}
