package member

import (
	"bbk/app/user/internal/dao"
	"bbk/app/user/internal/model"
	"bbk/app/user/internal/service/admin"
	"bbk/common/berror"
	"context"
)

type sMember struct{}

func init() {
	admin.RegisterMember(&sMember{})
}

func (s *sMember) GetMemberList(ctx context.Context, params *model.AdminGetMemberListDto) (res []model.AdminGetMemberListVo, total int, err error) {
	m := dao.AdminUser
	db := m.Ctx(ctx).Fields(res).Safe(false).Limit(params.GetOffset(), params.GetLimit())

	if params.Name != "" {
		name := "%" + params.Name + "%"
		db.Where(
			db.Builder().WhereLike(m.Columns().Name, name).
				WhereOr(m.Columns().Username, name),
		)
	}

	if params.UserID > 0 {
		db.Where(m.Columns().UserId, params.UserID)
	}

	if params.Status != nil {
		db.Where(m.Columns().Status, params.Status)
	}

	err = db.WithAll().ScanAndCount(&res, &total, false)
	if err != nil {
		err = berror.NewInternalError(err)
	}

	return
}

func (s *sMember) GetMember(ctx context.Context, userId int) (res *model.AdminGetMemberVo, err error) {
	m := dao.AdminUser
	err = m.Ctx(ctx).Fields(res).Where(m.Columns().UserId, userId).Scan(&res)
	if err != nil {
		err = berror.NewInternalError(err)
	}
	return
}
