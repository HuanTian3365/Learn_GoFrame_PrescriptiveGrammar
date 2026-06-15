package main

import (
	_ "goframe_prescriptivegrammar/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe_prescriptivegrammar/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
