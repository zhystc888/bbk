package controller

import (
	v1 "bbk/app/user/api/admin_member/v1"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

func init() {
	conn := grpcx.Client.MustNewGrpcClientConn("user")
	UserRpcService = v1.NewAdminMemberClient(conn)
}
