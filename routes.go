package main

import "net/http"

func Route(){
	http.HandleFunc("deal/", DealView)
	http.HandleFunc("claimDeal/", ClaimDealView)
}
