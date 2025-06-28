package error

import "github.com/gogf/gf/v2/errors/gcode"

var (
	CodeTokenIsEmpty = gcode.New(10401, "token不存在！", nil)
	CodeTokenInvalid = gcode.New(10401, "token验证失败！", nil)
	CodeTokenExpired = gcode.New(10401, "token已过期！", nil)
)

var (
	CodeInternal = gcode.New(10500, "服务器内部错误", nil)
)
