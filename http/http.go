package http

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"nginx-ldap-auth/g"
	"nginx-ldap-auth/http/controllers"
)

func Start() {
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "sessionID"
	beego.BConfig.WebConfig.EnableXSRF = true
	beego.BConfig.WebConfig.XSRFKey = "5ba12e0e0e34570a508ab27194d55075"
	beego.BConfig.WebConfig.XSRFExpire = g.Config().Http.XSRFExpire

	beego.SetStaticPath("/static", "static")

	if !g.Config().Http.Debug {
		logs.SetLevel(logs.LevelInformational)
	}
	ConfigRouters()
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run(g.Config().Http.Listen)
}
