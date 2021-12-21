package main

import (
	"time"
)

type SaleModelType struct {
	ID   int     `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time  `json:"end_time"`
	Products  []Item   `json:"products"`
}


func Save(saleModel SaleModelType) (int, error) {
	return 0, nil;
}

func Update(modelType SaleModelType) (int, error) {
	return 0, nil
}

func Delete(saleID int) (error) {
	return nil
}

func fetchSaleDetails(saleID int) (SaleModelType, error){
	return SaleModelType{}, nil;
}