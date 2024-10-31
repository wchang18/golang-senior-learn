package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	prfilev1 "user-admin/api/profile/v1"
	v1 "user-admin/api/user/v1"
	"user-admin/internal/consts"
	"user-admin/internal/model"
	"user-admin/internal/service"
)

type sUser struct{}

func init() {
	service.RegisterUser(NewUser())
}

func NewUser() *sUser {
	return new(sUser)
}

func (s *sUser) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	user, err := service.Store().GetUser(req.Username)
	if user.UserName == "" {
		err = gerror.New("username not exist")
		return
	}
	if user.Password != req.Password {
		err = gerror.New("password error")
		return
	}
	service.Session().Set(consts.UserNameKey, user.UserName)
	return
}

func (s *sUser) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	service.Session().Remove(consts.UserNameKey)
	return
}

func (s *sUser) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	one, _ := service.Store().GetUser(req.Username)
	if one.UserName == req.Username {
		err = gerror.New("username exist")
		return
	}
	user := model.User{
		UserName:  req.Username,
		Password:  req.Password,
		Telephone: req.Telephone,
	}
	err = service.Store().SetUser(user)
	return
}

func (l *sUser) GetUser(ctx context.Context, req *prfilev1.GetProfileReq) (res *prfilev1.GetProfileRes, err error) {
	name := ghttp.RequestFromCtx(ctx).GetCtxVar(consts.UserNameKey)
	user, err := service.Store().GetUser(name.String())
	res = &prfilev1.GetProfileRes{
		Username:  user.UserName,
		Telephone: user.Telephone,
	}
	return
}
