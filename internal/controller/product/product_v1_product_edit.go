package product

import (
	"context"
	"goframe_prescriptivegrammar/internal/service"

	"goframe_prescriptivegrammar/api/product/v1"
)

func (c *ControllerV1) ProductEdit(ctx context.Context, req *v1.ProductEditReq) (res *v1.ProductEditRes, err error) {
	data, err := service.Product().Edit(ctx, &req.ProductEditReq)
	if err != nil {
		return nil, err
	}
	return &v1.ProductEditRes{
		ProductEditRes: data,
	}, nil
}
