package product

import (
	"context"
	v1 "goframe_prescriptivegrammar/api/product/v1"
	"goframe_prescriptivegrammar/internal/service"
)

func (c *ControllerV1) Product(ctx context.Context, req *v1.ProductListReq) (res *v1.ProductListRes, err error) {
	data, err := service.Product().List(ctx, &req.ProductListReq)
	if err != nil {
		return nil, err
	}
	res = &v1.ProductListRes{
		ProductListRes: data,
	}
	return
}
