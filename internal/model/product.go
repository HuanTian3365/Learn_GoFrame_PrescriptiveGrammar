package model

type ProductListReq struct {
}

type ProductListRes struct {
	Items []string `json:"items"`
}
