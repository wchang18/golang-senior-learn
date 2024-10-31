package user

import (
	"context"
	"user-admin/internal/service"

	"user-admin/api/user/v1"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	return service.User().Register(ctx, req)
}
