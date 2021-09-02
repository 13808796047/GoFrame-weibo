package main

import (
	_ "GoFrame-weibo/boot"
	_ "GoFrame-weibo/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
