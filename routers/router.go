// @APIVersion 1.0.0
// @Title User center api
// @Description User center api
// @Contact zhaoyf12138@gmail.com
package routers

import (
	"usercenter/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/verification",
			beego.NSInclude(
				&controllers.VerificationController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.ErrorController(&controllers.ErrorController{})
}
