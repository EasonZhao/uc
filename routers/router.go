// @APIVersion 1.0.0
// @Title User center api
// @Description User center api
// @Contact zhaoyf12138@gmail.com
package routers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"strings"
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
				&controllers.AccountController{
					SecretKey: beego.AppConfig.String("secretkey"),
				},
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

	//filter
	beego.InsertFilter("/api/v1/account/info", beego.BeforeRouter, filterToken)
}

func filterToken(ctx *context.Context) {
	ss := ctx.Input.Header("Authorization")
	parseToken := func(str string) (string, error) {
		sv := strings.Split(str, " ")
		if len(sv) == 2 && strings.ToLower(sv[0]) == "bearer" && sv[1] != "" {
			return sv[1], nil
		}
		return "", errors.New("invalid authorization.")
	}
	tokenStr, err := parseToken(ss)
	if err != nil {
		result := controllers.NewRPCResult(controllers.STATUS_ERR)
		result.Data["code"] = err.Error()
		ctx.Output.JSON(result, true, false)
	}

	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("secretkey")), nil
	})
	if err != nil {
		result := controllers.NewRPCResult(controllers.STATUS_ERR)
		result.Data["code"] = err.Error()
		ctx.Output.JSON(result, true, false)
	}
	if !token.Valid {
		result := controllers.NewRPCResult(controllers.STATUS_ERR)
		result.Data["code"] = "invalid token"
		ctx.Output.JSON(result, true, false)
	}
	if v, err := strconv.Atoi(claims.Id); err == nil {
		ctx.Input.SetData("id", v)
	} else {
		//TODO log & return error
	}
}
