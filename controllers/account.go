package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"usercenter/models"
)

// UCS Verification API
type AccountController struct {
	beego.Controller
}

type registInfo struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	PhoneNum string `json:"phone"`
	Code     string `json:"code"`
}

//短信验证码确认
func checkEmailCode(email string, code string) bool {
	if code == "123456" {
		return true
	}
	return false
}

// @router /regist [post]
func (this *AccountController) Regist() {
	result := NewRPCResult(STATUS_OK)
	regitType := this.GetString("type")
	info := registInfo{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &info)
	for {
		if regitType == "email" {
			if !checkEmailCode(info.Email, info.Code) {
				result.Status = STATUS_ERR
				result.Data["code"] = "verification code error."
				break
			}
			u, err := models.RegistByEmail(info.Email, info.Password)
			if err != nil {
				result.Status = STATUS_ERR
				result.Data["code"] = err.Error()

			} else {
				result.Data["username"] = u.Username
			}
			break
		} else if regitType == "phone" {

		} else {
			result.Status = STATUS_ERR
			result.Data["code"] = "regist type not support"
			break
		}
	}
	this.Data["json"] = result
	this.ServeJSON()
}
