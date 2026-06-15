package service

import (
	"context"
	"goframe_prescriptivegrammar/internal/model"
)

type IProduct interface {
	List(ctx context.Context, req *model.ProductListReq) (res *model.ProductListRes, err error)
	View(ctx context.Context, req *model.ProductViewReq) (res *model.ProductViewRes, err error)
}

var product IProduct

func Product() IProduct {
	if product == nil {
		panic("product no init")
	}
	return product
}

func RegisterProduct(i IProduct) {
	product = i
}
