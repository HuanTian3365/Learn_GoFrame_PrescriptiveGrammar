// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShopProductDao is the data access object for the table shop_product.
type ShopProductDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ShopProductColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ShopProductColumns defines and stores column names for the table shop_product.
type ShopProductColumns struct {
	Id          string //
	ProductName string //
	ProductCode string //
	Price       string //
	Status      string //
	Remark      string //
	TenantId    string //
	CreatedBy   string //
	CreatedAt   string //
	UpdatedBy   string //
	UpdatedAt   string //
}

// shopProductColumns holds the columns for the table shop_product.
var shopProductColumns = ShopProductColumns{
	Id:          "id",
	ProductName: "product_name",
	ProductCode: "product_code",
	Price:       "price",
	Status:      "status",
	Remark:      "remark",
	TenantId:    "tenant_id",
	CreatedBy:   "created_by",
	CreatedAt:   "created_at",
	UpdatedBy:   "updated_by",
	UpdatedAt:   "updated_at",
}

// NewShopProductDao creates and returns a new DAO object for table data access.
func NewShopProductDao(handlers ...gdb.ModelHandler) *ShopProductDao {
	return &ShopProductDao{
		group:    "default",
		table:    "shop_product",
		columns:  shopProductColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ShopProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ShopProductDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ShopProductDao) Columns() ShopProductColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ShopProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ShopProductDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ShopProductDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
