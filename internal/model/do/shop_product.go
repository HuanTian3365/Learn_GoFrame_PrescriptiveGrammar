// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ShopProduct is the golang structure of table shop_product for DAO operations like Where/Data.
type ShopProduct struct {
	g.Meta      `orm:"table:shop_product, do:true"`
	Id          any         //
	ProductName any         //
	ProductCode any         //
	Price       any         //
	Status      any         //
	Remark      any         //
	TenantId    any         //
	CreatedBy   any         //
	CreatedAt   *gtime.Time //
	UpdatedBy   any         //
	UpdatedAt   *gtime.Time //
}
