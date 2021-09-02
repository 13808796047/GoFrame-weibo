package auth

import (
	"GoFrame-weibo/app/request"
	"GoFrame-weibo/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// 用户API管理对象
var Register = &registerApi{}

type registerApi struct{}


// Index is a demonstration route handler for output "Hello World!".
func (c *registerApi) ShowRegistrationForm(r *ghttp.Request) {

	r.Response.WriteTpl("layouts/app.html", g.Map{
		"content":    "auth/register.html",
		"title":      "注册",
		"page_class": "register-page",
	})
}
func (c *registerApi) Register(r *ghttp.Request) {

	// 1.验证
	var req *request.RegisterReq
	if err := r.Parse(&req); err != nil {
		if v, ok := err.(gvalid.Error); ok {
			r.Response.WriteTpl("layouts/app.html", g.Map{
				"content":    "auth/register.html",
				"title":      "注册",
				"page_class": "register-page",
				"user":&req,
				"errors":v.Maps(),
			})
		}
	}else{
		service.User.Create(r.Context(),req.Name,req.Email,req.Password)
	}
}
