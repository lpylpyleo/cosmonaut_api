package main

import (
	_ "cosmonaut_api/boot"
	_ "cosmonaut_api/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
