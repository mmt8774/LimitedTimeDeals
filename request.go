package main

import "time"

type DealRequestType struct {
	SaleID   int   `json:"sale_id"`
	Action  string `json:"action"`
	TotalQuantity int  `json:"total_quantity"`
	Products []Item `json:"products"`
	StartTime time.Time `json:"start_time"`
	EndTime   *time.Time  `json:"end_time"`
}

type ClaimDealRequestType struct {
	SaleID   int `json:"sale_id"`
	ProductID int `json:"product_id"`
	UserID    int  `json:"user_id"`
}
