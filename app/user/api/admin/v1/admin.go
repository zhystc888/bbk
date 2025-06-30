package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetAdminListReq struct {
	g.Meta `path:"/admin/list" tags:"getAdminList" method:"get" summary:"查看成员列表"`
	Name   string `json:"name"`
}

type GetAdminListRes struct {
	g.Meta `mime:"application/json"`
	Admin
	CreateAt string `json:"createAt"`
}

type GetAdminReq struct {
	g.Meta `path:"/admin" tags:"getAdmin" method:"get" summary:"查看成员详情"`
	ID     int `json:"id" v:"required"`
}

type GetAdminRes struct {
	g.Meta `mime:"application/json"`
	Admin
}

type Admin struct {
	UserID int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}
