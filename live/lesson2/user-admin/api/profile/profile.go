// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package profile

import (
	"context"

	"user-admin/api/profile/v1"
)

type IProfileV1 interface {
	GetProfile(ctx context.Context, req *v1.GetProfileReq) (res *v1.GetProfileRes, err error)
}
