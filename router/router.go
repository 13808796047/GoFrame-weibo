package router

import (
	"GoFrame-weibo/app/api"
	"GoFrame-weibo/app/api/auth"
	"GoFrame-weibo/app/middlewares"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middlewares.Ctx)
		group.GET("/", api.Home.Index)
		group.GET("/register", auth.Register.ShowRegistrationForm)
		group.POST("/register", auth.Register.Register)
		group.GET("/captcha", auth.Captcha.GenerateCaptchaHandler)
	})
}
