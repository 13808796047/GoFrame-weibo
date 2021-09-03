package auth

import (
	"GoFrame-weibo/app/dao"
	"GoFrame-weibo/app/helpers"
	"GoFrame-weibo/app/model"
	"GoFrame-weibo/app/service"
	"context"
	"github.com/gogf/gf/errors/gerror"
)

const (
	// 用户信息存放在Session中的Key
	sessionKeyUser = "SessionKeyUser"
)

// 获取登录用户
func User(ctx context.Context) *model.Users {
	customCtx := service.Context.Get(ctx)
	if customCtx != nil {
		if v := customCtx.Session.GetVar(sessionKeyUser); !v.IsNil() {
			var user *model.Users
			_ = v.Struct(&user)
			return user
		}
	}
	return nil
}

// Login登录指定用户
func Login(ctx context.Context, user *model.Users) error {
	return service.Context.Get(ctx).Session.Set(sessionKeyUser, user)
}

// 退出登录
func Logout(ctx context.Context) error {
	customCtx := service.Context.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyUser)
	}
	return nil
}

// Attempt 登录
func Attempt(ctx context.Context, email, password string) error {
	// 1. 根据Email获取用户
	user := model.Users{}
	err := dao.Users.Ctx(ctx).Where("email", email).Scan(&user)
	// 2. 如果出现错误
	if err != nil {
		return gerror.New("账号或密码不存在")
	}
	// 3.匹配密码
	if !helpers.CheckHash(password, user.Password) {
		return gerror.New("账号或密码不存在")
	}
	service.Context.Get(ctx).Session.Set(sessionKeyUser, &user)
	return nil
}
