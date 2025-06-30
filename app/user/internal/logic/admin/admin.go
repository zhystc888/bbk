package admin

import (
	"bbk/app/user/internal/dao"
	"bbk/app/user/internal/model"
	"bbk/app/user/internal/service"
	"bbk/common/berror"
	"context"
	"database/sql"
	"errors"
)

type sAdmin struct{}

func init() {
	service.RegisterAdmin(&sAdmin{})
}

func (s *sAdmin) GetAdminList(ctx context.Context) ([]model.Admin, error) {

}

func (s *sAdmin) GetAdmin(ctx context.Context, id int) (*model.Admin, error) {
	var admin *model.Admin
	err := dao.Admin.Ctx(ctx).Where("`id` = ?", id).Scan(&admin)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, berror.NewInternalError(err)
	}
	return admin, nil
}
