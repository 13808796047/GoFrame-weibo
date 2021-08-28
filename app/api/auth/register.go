package auth

import (
	"GoFrame-weibo/app/request"
	"GoFrame-weibo/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
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

	var req *request.RegisterReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteTpl("layouts/app.html", g.Map{
			"message":map[string]interface{}{
				"danger":err,
			},
			"title":      "提示",
		})
		r.Response.RedirectBack()
	}
	service.User.Create(r.Context())
}
