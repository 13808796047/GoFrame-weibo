package api

import (
	"GoFrame-weibo/app/request"
	"GoFrame-weibo/app/service"
	"GoFrame-weibo/library/upload"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
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
func (c *userApi) Edit(r *ghttp.Request) {
	id := r.GetString("id")
	user := service.User.Show(r.Context(), id)
	r.Response.WriteTpl("layouts/app.html", g.Map{
		"content":    "users/edit.html",
		"page_class": "users-page",
		"user":       user,
	})
}
func (c *userApi) Update(r *ghttp.Request) {
	var req *request.UserUpdateReq
	if err := r.Parse(&req); err != nil {
		if v, ok := err.(gvalid.Error); ok {
			r.Response.WriteTpl("layouts/app.html", g.Map{
				"content":    "users/edit.html",
				"title":      "用户编辑",
				"page_class": "register-page",
				"user":       &req,
				"errors":     v.Maps(),
			})
		}
	} else {
		file := r.GetUploadFile("avatar")
		filename, err := upload.Save(file, "avatar", r.GetString("id"))
		if err != nil {
			return
		}
		service.User.Update(r.Context(), r.GetString("id"), req.Name, req.Email, r.GetString("introduction"), filename)
		r.Response.RedirectTo("/users/show?id=" + r.GetString("id"))
	}
}
