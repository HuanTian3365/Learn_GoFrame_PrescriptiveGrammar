package response

import "github.com/gogf/gf/v2/net/ghttp"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 通用响应
func Json(r *ghttp.Request, code int, message string, data interface{}) {
	r.Response.WriteJson(Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// 先写响应,中断剩余流程,鉴权使用
func JsonExit(r *ghttp.Request, code int, message string, data interface{}) {
	Json(r, code, message, data)
	r.ExitAll()
}
