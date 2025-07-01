package v1

import "github.com/gogf/gf/v2/frame/g"

type AdminLoginReq struct {
	g.Meta `path:"/login" tags:"登录" method:"get" summary:"查看成员详情"`
}

type AdminLoginRes struct {
}
