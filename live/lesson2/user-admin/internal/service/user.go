// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	prfilev1 "user-admin/api/profile/v1"
	v1 "user-admin/api/user/v1"
)

type (
	IUser interface {
		Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
		Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
		Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error)
		GetUser(ctx context.Context, req *prfilev1.GetProfileReq) (res *prfilev1.GetProfileRes, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
