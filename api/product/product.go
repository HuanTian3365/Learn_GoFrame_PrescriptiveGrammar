// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package product

import (
	"context"

	"goframe_prescriptivegrammar/api/product/v1"
)

type IProductV1 interface {
	Product(ctx context.Context, req *v1.ProductListReq) (res *v1.ProductListRes, err error)
}
