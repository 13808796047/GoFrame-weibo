package middlewares

import (
	"GoFrame-weibo/app/model"
	"GoFrame-weibo/app/service"
	"github.com/gogf/gf/net/ghttp"
)

// 自定义上下文对象
func Ctx(r *ghttp.Request)  {
	// 初始化
	customCtx := &model.Context{
		Session: r.Session,
	}
	service.Context.Init(r,customCtx)
	r.Middleware.Next()
}