package middleware

import (
	"goframe_prescriptivegrammar/internal/response"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

func ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}

	err := r.GetError()
	res := r.GetHandlerResponse()

	if err != nil {
		response.Json(r, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.Json(r, 0, "", res)
}
