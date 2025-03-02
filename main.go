package main

import (
	_ "tree-table/internal/packed"

	_ "tree-table/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"tree-table/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
