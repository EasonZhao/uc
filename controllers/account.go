package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"usercenter/models"
)

const (
	TOKEN_EXP = 60 * 60 * 24
)

// UCS Verification API
type AccountController struct {
	SecretKey string
	beego.Controller
}

type registInfo struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	PhoneNum string `json:"phone"`
	Code     string `json:"code"`
}

type loginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &info); err != nil {
		result.Status = STATUS_ERR
		result.Data["code"] = "invalid param"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	for {
		if regitType == "email" {
			if !checkEmailCode(info.Email, info.Code) {
				result.Status = STATUS_ERR
				result.Data["code"] = "verification code error"
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

// @router /login [post]
func (this *AccountController) Login() {
	result := NewRPCResult(STATUS_OK)
	info := loginInfo{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &info); err != nil {
		result.Status = STATUS_ERR
		result.Data["code"] = "invalid param."
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	u, err := models.Login(info.Username, info.Password)
	if err != nil {
		result.Status = STATUS_ERR
		result.Data["code"] = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	//genrate token
	{
		signingKey := []byte(this.SecretKey)
		// Create the Claims
		sv, _ := beego.AppConfig.Int64("tokenexp")
		if sv <= 0 {
			sv = TOKEN_EXP
		}
		exprise := time.Now().Add(time.Second * time.Duration(sv)).Unix()
		claims := &jwt.StandardClaims{
			ExpiresAt: exprise,
			Issuer:    "usercenter",
			IssuedAt:  time.Now().Unix(),
			Id:        strconv.Itoa(u.Id),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(signingKey)
		if err != nil {
			logs.Error(err)
			result.Status = STATUS_ERR
			result.Data["code"] = "Error while signing the token"
			this.Data["json"] = result
			this.ServeJSON()
			return
		}
		result.Data["token"] = ss
		result.Data["exprise"] = exprise
		this.Data["json"] = result
		this.ServeJSON()
	}
}

// @router /info [get]
func (this *AccountController) Info() {
	result := NewRPCResult(STATUS_OK)
	id := this.Ctx.Input.GetData("id").(int)
	u, err := models.QueryUserById(id)
	if err != nil || u == nil {
		logs.Error("query user failure, id = ", id)
		result.Status = STATUS_ERR
		result.Data["code"] = "server internal error"
	} else {
		data := map[string]interface{}{
			"id":        u.Id,
			"username":  u.Username,
			"phone":     u.PhoneNum,
			"email":     u.Email,
			"authphone": u.AuthPhone,
			"authemail": u.AuthEmail,
		}
		result.Data = data
	}
	this.Data["json"] = result
	this.ServeJSON()
}
