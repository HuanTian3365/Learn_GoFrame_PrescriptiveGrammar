// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package product

import (
	"context"

	"goframe_prescriptivegrammar/api/product/v1"
)

type IProductV1 interface {
	ProductList(ctx context.Context, req *v1.ProductListReq) (res *v1.ProductListRes, err error)
	ProductView(ctx context.Context, req *v1.ProductViewReq) (res *v1.ProductViewRes, err error)
}
