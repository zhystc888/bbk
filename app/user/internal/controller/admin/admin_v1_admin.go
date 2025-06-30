package admin

import (
	"bbk/app/user/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"bbk/app/user/api/admin/v1"
)

func (c *ControllerV1) GetAdminList(ctx context.Context, req *v1.GetAdminListReq) (res *v1.GetAdminListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV1) GetAdmin(ctx context.Context, req *v1.GetAdminReq) (res *v1.GetAdminRes, err error) {
	admin, err := service.Admin().GetAdmin(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	if admin == nil {
		return nil, nil
	}
	res = &v1.GetAdminRes{}
	res.UserID = admin.UserID
	res.Name = admin.Name
	res.Status = admin.Status
	return res, nil
}
