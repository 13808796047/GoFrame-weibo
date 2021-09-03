package api

import (
	"GoFrame-weibo/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var User = &userApi{}

type userApi struct{}

func (c *userApi) Show(r *ghttp.Request) {
	id := r.GetString("id")
	user := service.User.Show(r.Context(), id)
	r.Response.WriteTpl("layouts/app.html", g.Map{
		"content":    "users/show.html",
		"page_class": "users-page",
		"user":       user,
	})
}
