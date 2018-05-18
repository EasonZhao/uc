package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/pquerna/otp/totp"
	"image/png"
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

type googleVerifyInfo struct {
	Code string `json:"code"`
}

var mc cache.Cache

func init() {
	var err error
	if mc, err = cache.NewCache("memory", `{"interval":60}`); err != nil {
		logs.Critical(err)
		panic(err)
	}
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

// @router /authotp [post]
func (this *AccountController) AuthOtp() {
	result := NewRPCResult(STATUS_OK)
	id := this.Ctx.Input.GetData("id").(int)
	u, err := models.QueryUserById(id)
	if err != nil || u == nil {
		logs.Error("query user failure, id = ", id)
		result.Status = STATUS_ERR
		result.Data["code"] = "server internal error"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	if u.GAAuth.Authed {
		result.Status = STATUS_ERR
		result.Data["code"] = "otp exist"
	} else {
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      "usercenter",
			AccountName: u.Username,
		})
		if err != nil {
			result.Status = STATUS_ERR
			result.Data["code"] = err.Error()
		} else {
			img, err := key.Image(520, 520)
			if err != nil {
				result.Status = STATUS_ERR
				result.Data["code"] = err.Error()
				this.Data["json"] = result
				this.ServeJSON()
			}
			var buf bytes.Buffer
			png.Encode(&buf, img)
			encodeStr := base64.StdEncoding.EncodeToString(buf.Bytes())
			result.Data["base64"] = encodeStr
			//put into cache
			name := "otp_" + u.Username
			mc.Put(name, key.Secret(), time.Minute*5)
			result.Data["code"] = key.Secret()
		}
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// @router /acceptotp [post]
func (this *AccountController) AcceptOtp() {
	result := NewRPCResult(STATUS_OK)
	id := this.Ctx.Input.GetData("id").(int)
	u, err := models.QueryUserById(id)
	if err != nil || u == nil {
		logs.Error("query user failure, id = ", id)
		result.Status = STATUS_ERR
		result.Data["code"] = "server internal error"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	info := struct {
		Code string `json:"code"`
	}{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &info); err != nil {
		result.Status = STATUS_ERR
		result.Data["code"] = "invalid param"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	if u.GAAuth.Authed {
		result.Status = STATUS_ERR
		result.Data["code"] = "otp exist"
	} else {
		name := "otp_" + u.Username
		val := mc.Get(name)
		if sercet, ok := val.(string); ok {
			if totp.Validate(info.Code, sercet) {
				mc.Delete(name)
				if err := u.AuthGA(sercet); err != nil {
					result.Status = STATUS_ERR
					result.Data["code"] = err.Error()
				}
			} else {
				result.Status = STATUS_ERR
				result.Data["code"] = "token mismatch"
			}
		}

	}
	this.Data["json"] = result
	this.ServeJSON()
}

// @router /otpverify [post]
func (this *AccountController) GoogleVerify() {
	result := NewRPCResult(STATUS_OK)
	info := struct {
		Code string `json:"code"`
	}{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &info); err != nil {
		result.Status = STATUS_ERR
		result.Data["code"] = "invalid param"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	id := this.Ctx.Input.GetData("id").(int)
	u, err := models.QueryUserById(id)
	if err != nil || u == nil {
		logs.Error("query user failure, id = ", id)
		result.Status = STATUS_ERR
		result.Data["code"] = "server internal error"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	if !u.GAAuth.Authed {
		result.Status = STATUS_ERR
		result.Data["code"] = "otp not auth"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	if totp.Validate(info.Code, u.GAAuth.Sercet) {
		this.Data["json"] = result
		this.ServeJSON()
		return
	} else {
		result.Status = STATUS_ERR
		result.Data["code"] = "token mismatch"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
}
