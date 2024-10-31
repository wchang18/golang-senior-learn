package user

import (
	"context"
	"user-admin/internal/service"

	"user-admin/api/user/v1"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	return service.User().Logout(ctx, req)
}
