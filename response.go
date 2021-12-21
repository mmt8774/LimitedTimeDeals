package main

type Response struct {
	Status   bool  `json:"status"`
	Code     int  `json:"code"`
	ErrMsg  string `json:"err_msg"`
}

type DealResponseType struct {
	Response
	DealResponse
}

type DealResponse struct {
	SaleID    int  `json:"sale_id"`
}
