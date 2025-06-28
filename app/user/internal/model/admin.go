package model

import "github.com/gogf/gf/v2/os/gtime"

type Admin struct {
	Id                uint64      `orm:"id,primary" json:"id" comment:"主键"`
	UserId            uint64      `orm:"user_id,unique" json:"user_id" comment:"用户ID"`
	Username          string      `orm:"username" json:"username" comment:"用户名"`
	Name              string      `orm:"name" json:"name" comment:"名称"`
	ResetPasswordTime *gtime.Time `orm:"reset_password_time" json:"reset_password_time" comment:"重置密码时间"`
	Status            uint8       `orm:"status" json:"status" comment:"状态：0未启用，1已启用，2禁止登陆"`
	Super             uint8       `orm:"super" json:"super" comment:"超级管理员，0否1是"`
	IsDeleted         uint8       `orm:"is_deleted" json:"is_deleted" comment:"数据状态0正常1删除"`
	CreateBy          uint64      `orm:"create_by" json:"create_by" comment:"创建人"`
	UpdateBy          uint64      `orm:"update_by" json:"update_by" comment:"修改人"`
	DeleteBy          uint64      `orm:"delete_by" json:"delete_by" comment:"删除人"`
	CreateAt          *gtime.Time `orm:"create_at" json:"create_at" comment:"创建时间"`
	UpdateAt          *gtime.Time `orm:"update_at" json:"update_at" comment:"更新时间"`
	DeletedAt         *gtime.Time `orm:"deleted_at" json:"deleted_at" comment:"删除时间"`
	User              *User       `orm:"with:user_id=id" json:"user" comment:"关联用户"`
}
