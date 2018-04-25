package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404() {
	result := NewRPCResult(STATUS_ERR)
	result.Data["code"] = "method don't exist"
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *ErrorController) Error500() {
	result := NewRPCResult(STATUS_ERR)
	result.Data["code"] = "server internal error"
	this.Data["json"] = result
	this.ServeJSON()
}
