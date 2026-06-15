// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShopProduct is the golang structure for table shop_product.
type ShopProduct struct {
	Id          int64       `json:"id"          orm:"id"           description:""` //
	ProductName string      `json:"productName" orm:"product_name" description:""` //
	ProductCode string      `json:"productCode" orm:"product_code" description:""` //
	Price       float64     `json:"price"       orm:"price"        description:""` //
	Status      int         `json:"status"      orm:"status"       description:""` //
	Remark      string      `json:"remark"      orm:"remark"       description:""` //
	TenantId    string      `json:"tenantId"    orm:"tenant_id"    description:""` //
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:""` //
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""` //
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"   description:""` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""` //
}
