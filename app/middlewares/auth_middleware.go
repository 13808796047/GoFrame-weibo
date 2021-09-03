package middlewares

import (
	"GoFrame-weibo/library/auth"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

func Auth(r *ghttp.Request) {
	if auth.User(r.Context()) != nil {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}
