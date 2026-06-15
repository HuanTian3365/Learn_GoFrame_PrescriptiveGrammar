package product

import (
	"context"
	"goframe_prescriptivegrammar/internal/dao"
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
	res = &model.ProductListRes{}
	err = dao.ShopProduct.Ctx(ctx).Scan(&res.Items)

	if err != nil {
		return nil, err
	}
	return res, nil
}
