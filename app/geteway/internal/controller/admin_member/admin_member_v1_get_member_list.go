package admin_member

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"bbk/app/geteway/api/admin_member/v1"
)

func (c *ControllerV1) GetMemberList(ctx context.Context, req *v1.GetMemberListReq) (res *v1.GetMemberListRes, err error) {
	
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
