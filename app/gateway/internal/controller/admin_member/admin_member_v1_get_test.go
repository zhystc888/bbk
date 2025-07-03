package admin_member

import (
	v1 "bbk/app/gateway/api/admin_member/v1"
	"context"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func TestControllerV1_Get(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := NewV1()
		ctx := context.Background()
		res, err := s.Get(ctx, &v1.GetReq{
			UserID: 1,
		})
	})
}
