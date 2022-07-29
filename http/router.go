package http

import (
	"nginx-ldap-auth/http/controllers"

	"github.com/astaxie/beego"
)

func ConfigRouters() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/auth-proxy", &controllers.AuthProxyController{})
	beego.Router("/api/v1/:control", &controllers.ControlController{})
}
