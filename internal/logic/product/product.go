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
	var items []*model.ProductListModel
	col := dao.ShopProduct.Columns()
	mod := dao.ShopProduct.Ctx(ctx)
	total := 0

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 产品码
	if req.ProductCode != "" {
		mod = mod.Where(col.ProductCode, req.ProductCode)
	}

	// 产品名模糊
	if req.ProductName != "" {
		mod = mod.WhereLike(col.ProductName, "%"+req.ProductName+"%")
	}

	// 按状态
	if req.Status != nil {
		mod = mod.Where(col.Status, req.Status)
	}

	err = mod.Fields(
		dao.ShopProduct.Columns().Id,
		dao.ShopProduct.Columns().ProductName,
		dao.ShopProduct.Columns().ProductCode,
		dao.ShopProduct.Columns().Price,
		dao.ShopProduct.Columns().Status,
	).Page(req.Page, req.PageSize).
		OrderDesc(col.Id).ScanAndCount(&items, &total, false)

	if err != nil {
		return nil, err
	}

	return &model.ProductListRes{
		Items:    items,
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
	}, nil
}
