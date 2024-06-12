package main

import (
	_ "gf_demo/internal/packed"

	_ "gf_demo/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"gf_demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
