package router

import (
	"goframe_prescriptivegrammar/internal/controller/health"
	"goframe_prescriptivegrammar/internal/controller/hello"
	"goframe_prescriptivegrammar/internal/controller/product"
	"goframe_prescriptivegrammar/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func InitRouter(s *ghttp.Server) {
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.ResponseHandler)
		group.Bind(
			hello.NewV1(),
		)
		group.Group("/api/v1", func(group *ghttp.RouterGroup) {
			group.Bind(
				health.NewV1(),
				product.NewV1(),
			)
		})
	})
}
