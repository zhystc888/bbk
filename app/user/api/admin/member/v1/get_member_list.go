package v1

import (
	"bbk/app/user/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type GetMemberListReq struct {
	g.Meta `path:"/member/list" tags:"成员" method:"get" sm:"查看成员列表"`
	model.AdminGetMemberListDto
	Page int `v:"min:0|分页号码错误" dc:"分页号码" d:"1"`
	Size int `v:"max:100#分页数量最大100条" dc:"分页数量，最大100" d:"10"`
}

type GetMemberListRes struct {
	g.Meta `mime:"application/json"`
	MemberBase
	LastTime string `json:"lastTime" dc:"最后登陆时间"`
	LastIp   string `json:"lastIp" dc:"最后登陆IP"`
	CreateAt string `json:"createAt" dc:"创建时间"`
}
