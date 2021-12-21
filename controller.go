package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func DealView(resp http.ResponseWriter, req *http.Request){
	var dealRequest DealRequestType
	var dealFinalResponse DealResponseType
	body,err := ioutil.ReadAll(req.Body)
	if err != nil {
		dealFinalResponse.Status = false
		dealFinalResponse.Code = http.StatusInternalServerError
		dealFinalResponse.ErrMsg = err.Error()
		return
	}
	err = json.Unmarshal(body, &dealRequest); if err != nil {
			logrus.Errorln("unable to unmarshal the request", err)
		dealFinalResponse.Status = false
		dealFinalResponse.Code = http.StatusInternalServerError
		dealFinalResponse.ErrMsg = err.Error()
	}

	dealResponse, err := DealAction(dealRequest); if err!= nil {
		dealFinalResponse.Status = true
		dealFinalResponse.Code = http.StatusInternalServerError
		dealFinalResponse.ErrMsg = err.Error()
	}

	if dealResponse != nil {
		dealFinalResponse.Status  = true
		dealFinalResponse.Code = http.StatusOK
		dealFinalResponse.SaleID = dealResponse.SaleID
	}
}


func ClaimDealView(resp http.ResponseWriter, req *http.Request){
	var claimDealRequest ClaimDealRequestType
	var dealFinalResponse DealResponseType
	body,err := ioutil.ReadAll(req.Body)
	if err != nil {
		dealFinalResponse.Status = false
		dealFinalResponse.Code = http.StatusInternalServerError
		dealFinalResponse.ErrMsg = err.Error()
		return
	}
	err = json.Unmarshal(body, &claimDealRequest); if err != nil {
		logrus.Errorln("ClaimDeal: unable to unmarshal the request", err)
		dealFinalResponse.Status = false
		dealFinalResponse.Code = http.StatusInternalServerError
		dealFinalResponse.ErrMsg = err.Error()
	}
	isValid := ValidateClaimDealRequest(claimDealRequest); if !isValid {
		dealFinalResponse.Status = false
		dealFinalResponse.Code = http.StatusInternalServerError
		dealFinalResponse.ErrMsg = "Invalid Request Type"
		return
	}

	ClaimDeal(claimDealRequest)

}

func ValidateClaimDealRequest(claimDealRequest ClaimDealRequestType) (bool){
	if claimDealRequest.SaleID <=0 || claimDealRequest.ProductID <=0 || claimDealRequest.UserID <= 0 {
		return false
	}
	return true
}