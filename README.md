#API Reference
##认证API
**GET /api/v1/verification/captcha 请求图片验证码**
请求参数：
|参数名称	|	是否必须	|类型	|描述	|默认值	|取值范围	|
|:-----	|:----:| :------: |:----:|:----:|:----|
|width	|false	|integer	|图片宽度	|240	|120-480
|height	|false	|integer	|图片高度	|80	|40-160
响应数据:
|参数名称	|	是否必须	|类型	|描述	|取值范围	|
|:-----	|:----:| :------: |:----|:----|
| status    |true	|string	|请求处理结果	|"ok","error"	|
|id	|true		|string	|验证码id	|	|
|base64-data	|true	|string	|图片base64值	|	|
|expiration	|true	|integer	|验证码有效期，单位秒	|600-1800	|
|width	|true	|integer	|图片宽度	|
|height	|true	|integer	|图片高度
data说明：
```json
{
	"status" : "ok",
	"data" : {
		"id" : "16WRvu4to1NGMsuNUHrP",
		"expiration" : 600,
		"base64" : "xxxxxxxxxxxxxxxx"，
		"width" : "240",
		"height" : "80"
	}
}
```
**POST /api/v1/verification/update 刷新图片验证码**
请求参数：
|参数名称	|	是否必须	|类型	|描述	|默认值	|取值范围	|
|:-----	|:----:| :------: |:----:|:----:|:----|
|id	|true	|string	|验证码id	|	
响应数据:
|参数名称	|	是否必须	|类型	|描述	|取值范围	|
|:-----	|:----:| :------: |:----|:----|
| status    |true	|string	|请求处理结果	|"ok","error"	|
|base64-data	|true	|string	|图片base64值	|	|
|expiration	|true	|integer	|验证码有效期，单位秒	|600-1800	|
data说明：
```json
{
	"status" : "ok",
	"data" : {
		"id" : "abcdefg1234567",
		"invalid" : 600,
		"base64-data" : "xxxxxxxxxxxxxxxx"
	}
}
```

**POST /api/v1/verification/verify 验证**
请求参数：
|参数名称	|	是否必须	|类型	|描述	|默认值	|取值范围	|
|:-----	|:----:| :------: |:----:|:----:|:----|
|id	|true	|string	|验证码id	
|digital	|true	|string	|内容
响应数据:
|参数名称	|	是否必须	|类型	|描述	|取值范围	|
|:-----	|:----:| :------: |:----|:----|
| status    |true	|string	|请求处理结果	|"ok","error"	|