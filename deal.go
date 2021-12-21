package main

import (
	"github.com/go-errors/errors"
	"time"
)

func DealAction(dealRequest DealRequestType) (*DealResponse,error){
	var dealResponse DealResponse
	switch (dealRequest.Action) {
	case ACTION_CREATE_DEAL:
		saleID, err := createDeal(dealRequest); if err != nil {
			return nil, errors.New("error while saving sale details")
		}
		dealResponse.SaleID = saleID
	case ACTION_DELETE_DEAL:
		if dealRequest.SaleID <= 0 {
			return nil, errors.New("invalid sale ID")
		}
		saleID ,err := deleteDeal(dealRequest); if err != nil {
			return nil, err
		}
		dealResponse.SaleID = saleID
	case ACTION_UPDATE_DEAL:
		if dealRequest.SaleID <= 0 {
			return nil, errors.New("invalid sale ID")
		}
		dealResponseTmp ,err := updateDeal(dealRequest); if err != nil {
		return nil, err
	}
		dealResponse = *dealResponseTmp
	default:
		return nil, errors.New("this action type is not defined")
	}
	return &dealResponse, nil
}



func createDeal(dealRequest DealRequestType) (int, error) {
	if dealRequest.StartTime.Unix() < time.Now().Unix() {
		return 0,  errors.New("invalid sale start time")
	}
	sale := SaleModelType{
		Products: dealRequest.Products,
		StartTime: dealRequest.StartTime,
		EndTime: *dealRequest.EndTime,
	}
	saleID, err := Save(sale); if err != nil {
		return 0, errors.New("createDeal: failed while saving")
	}
	return saleID, nil
}

func deleteDeal(dealRequest DealRequestType) (int, error){
	err := Delete(dealRequest.SaleID); if err != nil {
		return 0, errors.New("deleteDeal: failed while deleting")
	}
	return 0, nil
}

func updateDeal(dealRequest DealRequestType) (*DealResponse, error){
	saleDetails,err := fetchSaleDetails(dealRequest.SaleID); if err != nil {
		return nil, errors.New("error while fetching")
	}

	if dealRequest.EndTime != nil {
		saleDetails.EndTime = *dealRequest.EndTime
	}

	for i, item := range saleDetails.Products {
		for j, product := range dealRequest.Products {
			if item.Id == product.Id {
				saleDetails.Products[i].TotalQuantity = dealRequest.Products[j].TotalQuantity
			}
		}
	}

	saleID, err := Update(saleDetails); if err != nil {
		return nil, errors.New("error-while updating sale id")
	}
	dealResponse := &DealResponse{
		SaleID: saleID,
	}
	return dealResponse, nil
}

func claimDeal(claimDealView ClaimDealRequestType) (bool, error) {

}