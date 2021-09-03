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

		group.GET("/", api.Home.Index)
		group.Middleware(middlewares.Ctx)
		group.GET("/register", auth.Register.ShowRegistrationForm)
		group.POST("/register", auth.Register.Register)
		group.GET("/login", auth.Login.ShowLoginForm)
		group.POST("/login", auth.Login.Login)
		group.GET("/captcha", auth.Captcha.GenerateCaptchaHandler)
		group.POST("/logout", auth.Login.Logout)
		group.Middleware(middlewares.Auth)
		group.ALL("/users", api.User)
	})
}
