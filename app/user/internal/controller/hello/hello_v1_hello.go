package hello

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"runtime"

	"bbk/app/user/api/hello/v1"
)

func OpenFile() error {
	return gerror.New("permission denied")
}

func OpenConfig() error {
	return gerror.Wrap(OpenFile(), "configuration file opening failed")
}

func ReadConfig() error {
	return gerror.Wrap(OpenConfig(), "reading configuration failed")
}

type stack []uintptr

func callers(skip ...int) stack {
	var (
		pcs [64]uintptr
		n   = 3
	)
	if len(skip) > 0 {
		n += skip[0]
	}
	return pcs[:runtime.Callers(n, pcs[:])]
}

// Format 格式化 stack 为可读的调用栈
func (s stack) Format() string {
	if len(s) == 0 {
		return "no stack trace"
	}

	var result string
	frames := runtime.CallersFrames(s) // 将 PC 地址转换为 Frames
	for {
		frame, more := frames.Next()
		if frame.Function == "" {
			break // 无效帧
		}
		// 格式：函数名 文件名:行号
		result += fmt.Sprintf("%s\n    %s:%d\n", frame.Function, frame.File, frame.Line)
		if !more {
			break
		}
	}
	return result
}

func demoerror() error {
	err := gerror.NewCode(gcode.New(10020, "sql error", nil), "aaa", "bbb")
	code := gerror.Code(err)
	fmt.Printf("%+v\n", code)
	return err
}

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	//fmt.Printf("func: %s", callers().Format())
	//fmt.Printf("errconfig: %+v", ReadConfig())
	//fmt.Println(runtime.GOROOT())

	err = gerror.Wrap(demoerror(), "api calling failed")
	fmt.Printf("error: %+v", err)
	g.Log().Printf(ctx, "%+v", err)
	//res = nil
	return res, err
}
