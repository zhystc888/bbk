package member

import (
	"bbk/app/user/api/admin/member/v1"
	"bbk/app/user/internal/service/admin"
	"context"
)

func (c *ControllerV1) GetMemberList(ctx context.Context, req *v1.GetMemberListReq) (res v1.GetMemberListRes, err error) {
	list, total, err := admin.Member().GetMemberList(ctx, &req.AdminGetMemberListDto)
	if err != nil {
		return res, err
	}

	res = v1.GetMemberListRes{
		List:  []v1.GetMemberListItem{},
		Total: total,
	}

	for _, v := range list {
		item := v1.GetMemberListItem{}
		item.UserID = v.UserID
		item.Name = v.Name
		item.Status = v.Status
		item.CreateAt = v.CreateAt.String()
		item.LastIp = v.User.LastIp
		item.LastTime = v.User.LastTime.String()
		res.List = append(res.List, item)
	}

	return res, nil
}
