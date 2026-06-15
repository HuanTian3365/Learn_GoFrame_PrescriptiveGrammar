package product

import (
	"context"
	"goframe_prescriptivegrammar/internal/model"
	"goframe_prescriptivegrammar/internal/service"
)

type sProduct struct{}

func CreateProduct() *sProduct {
	return &sProduct{}
}

func init() {
	service.RegisterProduct(CreateProduct())
}

func (s *sProduct) List(ctx context.Context, req *model.ProductListReq) (res *model.ProductListRes, err error) {
	return &model.ProductListRes{
		Items: []string{"111", "222", "333"},
	}, nil
}
