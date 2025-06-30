package member

import (
	"bbk/app/user/api/admin/member/v1"
	"bbk/app/user/internal/service/admin"
	"context"
)

func (c *ControllerV1) GetMemberList(ctx context.Context, req *v1.GetMemberListReq) (res []v1.GetMemberListRes, err error) {
	list, err := admin.Member().GetMemberList(ctx, &req.AdminGetMemberListDto)
	if err != nil {
		return nil, err
	}

	for _, v := range list {
		item := v1.GetMemberListRes{}
		item.UserID = v.UserID
		item.Name = v.Name
		item.Status = v.Status
		item.CreateAt = v.CreateAt.String()
		item.LastIp = v.User.LastIp
		item.LastTime = v.User.LastTime.String()
		res = append(res, item)
	}

	return res, nil
}
