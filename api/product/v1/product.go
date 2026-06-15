package v1

import (
	"goframe_prescriptivegrammar/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ProductListReq struct {
	g.Meta `path:"/product/list" method:"get"`
	model.ProductListReq
}

type ProductListRes struct {
	*model.ProductListRes
}
