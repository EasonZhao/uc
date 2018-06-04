# API Reference
## 认证API
**1、 GET /api/v1/verification/captcha 请求图片验证码**

请求参数：

|参数名称	|	是否必须	|类型	|描述	|默认值	|取值范围	|
|:-----	|:----:| :------: |:----:|:----:|:----|
|width	|false	|integer	|图片宽度	|240	|120-480|
|height	|false	|integer	|图片高度	|80	|40-160|

响应数据:

|参数名称	|	是否必须	|类型	|描述	|取值范围	|
|:-----	|:----:| :------: |:----|:----|
| status    |true	|string	|请求处理结果	|"ok","error"	|
|id	|true		|string	|验证码id	|	|
|base64-data	|true	|string	|图片base64值	|	|
|expiration	|true	|integer	|验证码有效期，单位秒	|600-1800	|
|width	|true	|integer	|图片宽度	|	|
|height	|true	|integer	|图片高度	|	|

data说明：
```json
{
	"status": "ok",
	"data": {
		"id": "16WRvu4to1NGMsuNUHrP",
		"expiration": 600,
		"base64": "xxxxxxxxxxxxxxxx",
		"width": 240,
		"height": 80
	}
}
```

## 获取用户信息API
**2、GET /api/v1/account/info 获取用户信息**

请求参数：

|参数名称	|	是否必须	|类型	|描述	|默认值	|取值范围	|
|:-----	|:----:| :------: |:----:|:----:|:----|
|id		|true	|string		|用户id	|	|	|

响应数据:

|参数名称	|	是否必须	|类型	|描述	|取值范围	|
|:-----	|:----:| :------: |:----|:----|
|status    |true	|string	|请求处理结果	|"ok","error"	|
|id	|true		|int	|用户id	|		|
|username	|true		|string	|用户名称	|		|
|phone	|true		|string	|用户手机号	|		|
|email	|true		|string	|用户email	|		|
|authphone	|true		|string	|验证手机	|		|
|authemail	|true		|string	|验证email	|		|	

data说明：
```json
{
	"status": "ok",
	"data":  {
		"id": "98720",
		"username": "xuefeng",
		"phone": "18918885678",
		"email": "andy@163.com",
		"authphone": "18918885679",
		"authemail": "www@163.com"
	}
}
```

## Email验证API
**3、GET /api/v1/email/verification Email验证**

请求参数：

|参数名称	|	是否必须	|类型	|描述	|默认值	|取值范围	|
|:-----	|:----:| :------: |:----:|:----:|:----|
|email		|true	|string	|用户email地址	|	|	|

响应数据:

|参数名称	|	是否必须	|类型	|描述	|取值范围	|
|:-----	|:----:| :------: |:----|:----|
|status    |true	|string	|请求处理结果	|"ok","error"	|
|email	|true		|string	|用户email地址	|		|
|expiration	|true	|int	|到期时间	|		|

data说明：
```json
{
	"status": "ok",
	"data":  {
		"email": "andy@163.com",
		"expiration": "1407408983"
	}
}
```

## google Opt认证API
**4、POST /api/v1/account/acceptotp google Opt认证API**

请求参数：

|参数名称	|	是否必须	|类型	|描述	|默认值	|取值范围	|
|:-----	|:----:| :------: |:----:|:----:|:----|
|id		|true	|string		|用户id	|	|	|
|code		|true	|string		|google opt码	|	|	|

响应数据:

|参数名称	|	是否必须	|类型	|描述	|取值范围	|
|:-----	|:----:| :------: |:----|:----|
|status    |true	|string	|请求处理结果	|"ok","error"	|

data说明：
```json
{
	"status": "ok"
}
```

##  生成认证二维码API
**5、POST /api/v1/account/authotp 生成认证二维码**

请求参数：

|参数名称	|	是否必须	|类型	|描述	|默认值	|取值范围	|
|:-----	|:----:| :------: |:----:|:----:|:----|
|Authorization		|true	|string		|Bearer token	|	|	|

响应数据:

|参数名称	|	是否必须	|类型	|描述	|取值范围	|
|:-----	|:----:| :------: |:----|:----|
|status    |true	|string	|请求处理结果	|"ok","error"	|
|base64    |true	|string	|Google验证器扫描的二维码	| 	|
|code    |true	|string	|在Google验证器中添加该密钥	|	|

data说明：
```json
{
	"status": "ok",
	"data":  {
		"base64": "YXNkZmFzZmFzZGZhc2ZyeWxrbmN3MjU0NTY3NmhnY3pzYWVmZ2hqamxlZXJ0ZXR2Y2dma3JhYQ== ",
		"code": "HLHFO552WVT3UTB2 "
	}
}
```

##  登录API
**6、POST /api/v1/account/login 登录**

请求参数：

|参数名称	|	是否必须	|类型	|描述	|默认值	|取值范围	|
|:-----	|:----:| :------: |:----:|:----:|:----|
|username		|true	|string		|用户名	|	|	|
|password		|true	|string		|用户密码	|	|	|

响应数据:

|参数名称	|	是否必须	|类型	|描述	|取值范围	|
|:-----	|:----:| :------: |:----|:----|
|status    |true	|string	|请求处理结果	|"ok","error"	|
|token    |true	|string	|用户token	| 	|
|exprise    |true	|int	|过期时间	|	|

data说明：
```json
{
	"status": "ok",
	"data":  {
		"token": "YXNkZmFzZmFzZGZhc2Zye",
		"exprise": "2018-05-22 12:00:00"
	}
}
```

##  google认证API
**7、POST /api/v1/account/otpverify google认证**

请求参数：

|参数名称	|	是否必须	|类型	|描述	|默认值	|取值范围	|
|:-----	|:----:| :------: |:----:|:----:|:----|
|Authorization	|true	|string		|Bearer token	|	|	|
|code		|true	|string		|google opt码	|	|	|

响应数据:

|参数名称	|	是否必须	|类型	|描述	|取值范围	|
|:-----	|:----:| :------: |:----|:----|
|status    |true	|string	|请求处理结果	|"ok","error"	|

data说明：
```json
{
	"status": "ok",
}
```

##  注册API
**8、POST /api/v1/account/regist 注册**

请求参数：

|参数名称	|	是否必须	|类型	|描述	|默认值	|取值范围	|
|:-----	|:----:| :------: |:----:|:----:|:----|
|type		|true	|string		|注册类型	|	|	|
|email		|true	|string		|email	|	|	|
|password		|true	|string		|密码	|	|	|
|phone		|true	|string		|手机号码	|	|	|
|code		|true	|string		|验证码	|	|	|

响应数据:

|参数名称	|	是否必须	|类型	|描述	|取值范围	|
|:-----	|:----:| :------: |:----|:----|
|status    |true	|string	|请求处理结果	|"ok","error"	|
|username    |true	|string	|用户名	| 	|

data说明：
```json
{
	"status": "ok",
	"data":  {
		"username": "andy"
	}
}
```
