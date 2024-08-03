package BillingService

import (
	BillingRepository "pos-api/api/billing/repository"
	Schema "pos-api/api/billing/schema"
)

func CreateBill(body Schema.CreateBillDto) error {
	ID, err := BillingRepository.CreateBill(body.BillTotalPrice)
	if err != nil {
		return err
	}
	err = BillingRepository.CreateBillDetail(body.BillDetail, ID)
	if err != nil {
		return err
	}
	return nil
}
