// @APIVersion 1.0.0
// @Title cyber-life Auth API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact dev@cyber-life.cn
// @TermsOfServiceUrl http://cyber-life.cn/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"cyber-auth-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api/auth",
		beego.NSNamespace("/register",
			beego.NSInclude(
				&controllers.RegisterController{},
			),
		),
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/ticket",
			beego.NSInclude(
				&controllers.TicketController{},
			),
		),
		beego.NSNamespace("/register/code",
			beego.NSInclude(
				&controllers.RegisterCodeController{},
			),
		),
		beego.NSNamespace("/lostpwd/code",
			beego.NSInclude(
				&controllers.LostpwdCodeController{},
			),
		),
		beego.NSNamespace("/lostpwd",
			beego.NSInclude(
				&controllers.LostpwdController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
