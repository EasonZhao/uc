// @APIVersion 1.0.0
// @Title User center api
// @Description User center api
// @Contact zhaoyf12138@gmail.com
package routers

import (
	"github.com/astaxie/beego"
	"usercenter/controllers"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/verification",
			beego.NSInclude(
				&controllers.VerificationController{},
			),
		),
		beego.NSNamespace("/account",
			beego.NSInclude(
				&controllers.AccountController{},
			),
		),
		beego.NSNamespace("/email",
			beego.NSInclude(
				&controllers.EmailController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.ErrorController(&controllers.ErrorController{})
}
