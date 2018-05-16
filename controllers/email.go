package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"usercenter/util"
)

type EmailController struct {
	beego.Controller
}

// @router /verification [get]
func (this *EmailController) Verification() {
	email := this.GetString("email")
	result := NewRPCResult(STATUS_OK)
	if !util.CheckEmail(email) {
		result.Status = STATUS_ERR
		result.Data["code"] = "invalid email"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.Data["email"] = email
	exp_t := time.Now().UTC().Unix() + int64(60*30)
	result.Data["expiration"] = strconv.FormatInt(exp_t, 10)
	this.Data["json"] = result
	this.ServeJSON()
}
