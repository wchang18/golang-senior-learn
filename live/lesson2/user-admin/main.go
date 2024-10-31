package main

import (
	_ "user-admin/internal/packed"

	_ "user-admin/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"user-admin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
