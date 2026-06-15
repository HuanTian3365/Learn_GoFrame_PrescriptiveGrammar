package model

type ProductListModel struct {
	Id          int64   `json:"id"`
	ProductName string  `json:"productName"`
	ProductCode string  `json:"productCode"`
	Price       float64 `json:"price"`
	Status      int     `json:"status"`
}

type ProductListReq struct {
}

type ProductListRes struct {
	Items []*ProductListModel `json:"items"`
}
