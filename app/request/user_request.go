package request

import (
	"GoFrame-weibo/app/model"
)

type User *model.Users
type RegisterReq struct {
	Name                 string `p:"name"  v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Email                string `p:"email"  v:"required|not_exists:users,email"`
	Password             string `p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
	PasswordConfirmation string `p:"password_confirmation" v:"required|length:6,30|same:password#请确认密码|密码长度不够|两次密码不一致"`
	Captcha              string `p:"captcha" v:"required|check-captcha#请输入验证码|请输入正确验证码"`
}

type UserUpdateReq struct {
	Id           string
	Name         string `p:"name"  v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Email        string `p:"email"  v:"required|not_exists:users,email"`
	Introduction string
	Avatar       string
}
