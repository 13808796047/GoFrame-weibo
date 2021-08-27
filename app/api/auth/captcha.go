package auth

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/steambap/captcha"
)

// 图形验证码
var Captcha = &captchaApi{}

type captchaApi struct{}

func (c *captchaApi) GenerateCaptchaHandler(r *ghttp.Request) {
	img, _ := captcha.New(90, 35)
	r.Session.Set("captcha", img.Text)
	img.WriteImage(r.Response.Writer)
}
