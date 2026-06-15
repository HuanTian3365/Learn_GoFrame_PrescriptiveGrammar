package model

type ProductListModel struct {
	Id          int64   `json:"id"`
	ProductName string  `json:"productName"`
	ProductCode string  `json:"productCode"`
	Price       float64 `json:"price"`
	Status      int     `json:"status"`
}

type ProductListReq struct {
	Page        int    `json:"page" d:"1"`
	PageSize    int    `json:"pageSize" d:"10"`
	ProductName string `json:"productName"`
	ProductCode string `json:"productCode"`
	Status      *int   `json:"status"`
}

type ProductListRes struct {
	Items    []*ProductListModel `json:"items"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"pageSize"`
	Total    int                 `json:"total"`
}
