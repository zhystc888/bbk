package model

import (
	v1 "bbk/app/user/api/admin/v1"
	"github.com/gogf/gf/v2/os/gtime"
)

type AdminListItem struct {
	v1.Admin
	CreateAt *gtime.Time
}

type Admin struct {
	v1.Admin
}
