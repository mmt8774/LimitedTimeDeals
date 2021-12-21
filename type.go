package main

import "time"

type User struct {
	Id    int64  `json:"user"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Item struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	TotalQuantity int `json:"total_quantity"`
	SoldQuantity int         `json:"sold_quantity"`
}

type Sale struct {
	Products  map[Item]int //map of item and total quantity
	StartTime time.Time    `json:"start_time"`
	EndTime   time.Time    `json:"end_time"`
}
