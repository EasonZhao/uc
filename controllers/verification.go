package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"strconv"
	"time"
)

const (
	WIDTH  = 240
	HEIGHT = 80
	LENGTH = 6
)

// UCS Verification API
type VerificationController struct {
	beego.Controller
}

// @Title Verification captcha
// @Description Request captcha image
// @Success 200 {status: "ok", data:
// @Param width query int false "image width"
// @Param height query int false "image height"
//	{id:"36c6kqY15iMGawLDi82R", base64-data:"data string", width:240, height:80, valid-time:600}}
// @Failure 500 server internal error
// @router /captcha [get]
func (this *VerificationController) Captcha() {
	width, _ := this.GetInt("width")
	if width < 120 || width > 480 {
		width = WIDTH
	}
	height, _ := this.GetInt("height")
	if height < 40 || height > 160 {
		height = HEIGHT
	}
	id := captcha.NewLen(LENGTH)
	digits := captcha.RandomDigits(LENGTH)
	image := captcha.NewImage(id, digits, width, height)
	b := bytes.NewBuffer(make([]byte, 0))
	_, err := image.WriteTo(b)
	if err != nil {
		beego.Error("captcha image write err=", err)
		this.Abort("500")
	}
	encode_str := base64.StdEncoding.EncodeToString(b.Bytes())
	result := NewRPCResult(STATUS_OK)
	result.Data["id"] = id
	result.Data["width"] = strconv.Itoa(width)
	result.Data["height"] = strconv.Itoa(height)
	result.Data["base64"] = encode_str
	exp_t := time.Now().UTC().Unix() + int64(captcha.Expiration.Seconds())
	result.Data["expiration"] = strconv.FormatInt(exp_t, 10)
	this.Data["json"] = result

	this.ServeJSON()
}
