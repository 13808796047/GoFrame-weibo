package request

import (
	"GoFrame-weibo/app/dao"
	"GoFrame-weibo/app/service"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gvalid"
	"strings"
)

func init()  {
	registerRule := "check-captcha"
	gvalid.RegisterRule(registerRule, func(ctx context.Context, rule string, value interface{}, message string, data interface{}) error {
		val := value.(string)
		if val != service.Context.Get(ctx).Session.Get("captcha") {
			return errors.New("验证码不正确")
		}
		return nil
	})
	// not_exists:users,email
	gvalid.RegisterRule("not_exists", func(ctx context.Context, rule string, value interface{}, message string, data interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")
		//tableName := rng[0]
		dbFiled := rng[1]
		val := value.(string)

		//model.DB.Table(tableName).Where(dbFiled+" = ?", val).Count(&count)
		r,_ := dao.Users.Ctx(ctx).Where(dbFiled, val).Count()
		g.Dump(r)

		if r != 0 {

			return errors.New(fmt.Sprintf("%v 已被占用", val))
		}
		return nil
	})
}