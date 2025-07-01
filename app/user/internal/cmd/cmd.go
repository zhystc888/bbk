package cmd

import (
	"bbk/app/user/internal/controller/admin/member"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

type JsonRes struct {
	Code    int         `json:"code" x-apifox-mock:"0"`     // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message" x-apifox-mock:"OK"` // 提示信息
	Data    interface{} `json:"data"`                       // 返回数据(业务接口定义具体数据结构)
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				s   = g.Server()
				oai = s.GetOpenApi()
			)

			// OpenApi自定义信息
			oai.Info.Title = `版本库`
			oai.Info.Description = `传奇版本库`
			oai.Config.CommonResponse = JsonRes{
				Code:    0,
				Message: "OK",
			}
			oai.Config.CommonResponseDataField = `Data`

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(func(r *ghttp.Request) {
					r.Response.CORSDefault()
					r.Middleware.Next()
				})
				group.Bind(
					member.NewV1(),
				)
			})
			//openApi(s)
			s.Run()
			return nil
		},
	}
)

func openApi(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info.Title = `版本库`
	openapi.Info.Description = `传奇版本库`
}
