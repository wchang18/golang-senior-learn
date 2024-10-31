package user

import (
	"context"
	"user-admin/internal/service"

	"user-admin/api/user/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	return service.User().Login(ctx, req)
}
