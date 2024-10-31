package profile

import (
	"context"
	"user-admin/internal/service"

	"user-admin/api/profile/v1"
)

func (c *ControllerV1) GetProfile(ctx context.Context, req *v1.GetProfileReq) (res *v1.GetProfileRes, err error) {
	return service.User().GetUser(ctx, req)
}
