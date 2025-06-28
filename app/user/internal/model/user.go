package model

import "github.com/gogf/gf/v2/os/gtime"

type User struct {
	Id        uint64      `orm:"id,primary" json:"id" comment:"主键"`
	Group     uint8       `orm:"group" json:"group" comment:"用户组 1. 后台管理员 2. 陪玩端 3. 客户端"`
	Password  string      `orm:"password" json:"-" comment:"密码"`
	LastTime  uint64      `orm:"last_time" json:"last_time" comment:"上次登陆时间"`
	LastIp    string      `orm:"last_ip" json:"last_ip" comment:"上次登陆ip"`
	IsDeleted uint8       `orm:"is_deleted" json:"is_deleted" comment:"数据状态0正常1删除"`
	CreateBy  uint64      `orm:"create_by" json:"create_by" comment:"创建人"`
	UpdateBy  uint64      `orm:"update_by" json:"update_by" comment:"修改人"`
	DeleteBy  uint64      `orm:"delete_by" json:"delete_by" comment:"删除人"`
	CreateAt  *gtime.Time `orm:"create_at" json:"create_at" comment:"创建时间"`
	UpdateAt  *gtime.Time `orm:"update_at" json:"update_at" comment:"更新时间"`
	DeletedAt *gtime.Time `orm:"deleted_at" json:"deleted_at" comment:"删除时间"`
}
