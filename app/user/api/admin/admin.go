// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"bbk/app/user/api/admin/v1"
)

type IAdminV1 interface {
	GetAdminList(ctx context.Context, req *v1.GetAdminListReq) (res *v1.GetAdminListRes, err error)
	GetAdmin(ctx context.Context, req *v1.GetAdminReq) (res *v1.GetAdminRes, err error)
}
