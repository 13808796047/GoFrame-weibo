package model

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	// 上下文变量存储键名,前后端系统共享
	ContextKey = "ContextKey"
)

// 请求上下文结构

type Context struct {
	Session *ghttp.Session
	User *ContextUser
	Data *g.Map
}

// 上下文中的用户信息
type ContextUser struct {
	Id uint
	Name string
	Email string
	Avatar string
}
