package auth

import (
	"GoFrame-weibo/app/service"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"github.com/steambap/captcha"
)

// 图形验证码
var Captcha = &captchaApi{}

type captchaApi struct{}

func (c *captchaApi) GenerateCaptchaHandler(r *ghttp.Request) {

	img, _ := captcha.NewMathExpr(90, 35)
	service.Context.Get(r.Context()).Session.Set("captcha",gstr.ToLower(img.Text))

	//r.Session.Set("captcha", img.Text)
	img.WriteImage(r.Response.Writer)
}
