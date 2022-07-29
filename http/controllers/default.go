package controllers

import (
	"nginx-ldap-auth/g"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	uname := this.GetSession("uname")
	if uname == nil {
		this.Ctx.Redirect(302, "/login")
		return
	}
	this.Ctx.Output.Body([]byte("nginx-ldap-auth, version " + g.VERSION))
}
