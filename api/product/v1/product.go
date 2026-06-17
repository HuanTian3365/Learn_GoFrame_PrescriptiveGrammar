package v1

import (
	"goframe_prescriptivegrammar/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 列表 ====================
type ProductListReq struct {
	g.Meta `path:"/product/list" method:"get"`
	model.ProductListReq
}

type ProductListRes struct {
	*model.ProductListRes
}

// ==================== 详情 ====================
type ProductViewReq struct {
	g.Meta `path:"/product/view" method:"get"`
	model.ProductViewReq
}

type ProductViewRes struct {
	*model.ProductViewRes
}

// ==================== 编辑 ====================
type ProductEditReq struct {
	g.Meta `path:"/product/edit" method:"post"`
	model.ProductEditReq
}

type ProductEditRes struct {
	*model.ProductEditRes
}
