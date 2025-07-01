package main

import (
	_ "bbk/app/geteway/internal/packed"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/os/gctx"

	"bbk/app/geteway/internal/cmd"
)

func main() {
	// 设置进程全局时区
	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	
	cmd.Main.Run(gctx.GetInitCtx())
}
