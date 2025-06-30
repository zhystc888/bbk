package member

import (
	"bbk/app/user/api/admin/member/v1"
	"bbk/app/user/internal/service/admin"
	"context"
)

func (c *ControllerV1) GetMember(ctx context.Context, req *v1.GetMemberReq) (res *v1.GetMemberRes, err error) {
	member, err := admin.Member().GetMember(ctx, req.UserID)
	if member != nil {
		res = &v1.GetMemberRes{}
		res.Name = member.Name
		res.UserID = member.UserID
		res.Status = member.Status
		res.Username = member.Username
	}

	return
}
