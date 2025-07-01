// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package admin

import (
	"bbk/app/user/internal/model"
	"context"
)

type (
	IMember interface {
		GetMemberList(ctx context.Context, params *model.AdminGetMemberListDto) (res []model.AdminGetMemberListVo, total int, err error)
		GetMember(ctx context.Context, userId int) (res *model.AdminGetMemberVo, err error)
	}
)

var (
	localMember IMember
)

func Member() IMember {
	if localMember == nil {
		panic("implement not found for interface IMember, forgot register?")
	}
	return localMember
}

func RegisterMember(i IMember) {
	localMember = i
}
