package auth

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 用户API管理对象
var Register = &registerApi{}

type registerApi struct{}
type RegisterReq struct {
	Name                 string `p:"name"  v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Email                string `p:"email"  v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Password             string `p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
	PasswordConfirmation string `p:"password_confirmation" v:"required|length:6,30|same:password#请确认密码|密码长度不够|两次密码不一致"`
	Captcha              string `p:"captcha" v:"required"`
}

// Index is a demonstration route handler for output "Hello World!".
func (c *registerApi) ShowRegistrationForm(r *ghttp.Request) {
	users,_ := g.Model("user").All()
	fmt.Println(users)
	r.Response.WriteTpl("layouts/app.html", g.Map{
		"content":    "auth/register.html",
		"title":      "注册",
		"page_class": "register-page",
	})
}
func (c *registerApi) Register(r *ghttp.Request) {

}
