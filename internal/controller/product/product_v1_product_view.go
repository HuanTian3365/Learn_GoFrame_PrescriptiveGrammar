package product

import (
	"context"
	"goframe_prescriptivegrammar/internal/service"

	"goframe_prescriptivegrammar/api/product/v1"
)

func (c *ControllerV1) ProductView(ctx context.Context, req *v1.ProductViewReq) (res *v1.ProductViewRes, err error) {
	data, err := service.Product().View(ctx, &req.ProductViewReq)
	if err != nil {
		return nil, err
	}
	return &v1.ProductViewRes{
		ProductViewRes: data,
	}, nil
}
