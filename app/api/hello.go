package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 用户API管理对象
var Home = &homeApi{}

type homeApi struct{}

// Index is a demonstration route handler for output "Hello World!".
func (c *homeApi) Index(r *ghttp.Request) {
	r.Response.WriteTpl("layouts/app.html", g.Map{
		"content": "pages/home.html",
		"title":   "首页",
	})
}
