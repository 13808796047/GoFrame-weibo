package auth

import (
	"GoFrame-weibo/library/auth"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

var Login = &loginApi{}

type loginApi struct{}

func (l *loginApi) ShowLoginForm(r *ghttp.Request) {
	r.Response.WriteTpl("layouts/app.html", g.Map{
		"content":    "auth/login.html",
		"title":      "登录",
		"page_class": "login-page",
	})
}

func (l *loginApi) Login(r *ghttp.Request) {
	// 1.初始化表单
	email := r.GetString("email")
	password := r.GetString("password")
	if err := auth.Attempt(r.Context(), email, password); err == nil {
		r.Response.RedirectTo("/", http.StatusFound)
	} else {
		r.Response.WriteTpl("layouts/app.html", g.Map{
			"content":    "auth/login.html",
			"title":      "登录",
			"page_class": "login-page",
			"email":      email,
			"password":   password,
			"errors":     err.Error(),
		})
	}
}
func (l loginApi) Logout(r *ghttp.Request) {
	auth.Logout(r.Context())
	r.Response.RedirectTo("/")
}
