// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package member

import (
	"context"

	"bbk/app/user/api/admin/member/v1"
)

type IMemberV1 interface {
	GetMember(ctx context.Context, req *v1.GetMemberReq) (res *v1.GetMemberRes, err error)
	GetMemberList(ctx context.Context, req *v1.GetMemberListReq) (res v1.GetMemberListRes, err error)
}
