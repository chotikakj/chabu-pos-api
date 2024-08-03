package ReportService

import (
	ReportRepository "pos-api/api/report/repository"
	Schema "pos-api/api/report/schema"
)

func GetHomeReport(body Schema.ReportHomeDto) (*Schema.ReportHome, error) {
	bill, err := ReportRepository.GetBillReport(body.StartDate, body.EndDate)
	if err != nil {
		return nil, err
	}
	billDetail, err := ReportRepository.GetBillDetailReport(body.StartDate, body.EndDate)
	if err != nil {
		return nil, err
	}
	if len(bill) == 0 {
		return nil, err
	}
	total_price := 0
	total_profit := 0
	order_amount := 0
	var best_sells []Schema.BestSellItems
	for _, val := range billDetail {
		total_price += val.BillDetailItemTotalPrice
		total_profit += val.BillDetailItemTotalPrice - val.BillDetailItemTotalCost
		best_sell := Schema.BestSellItems{
			Name:        val.BillDetailItemName,
			Pricing:     val.BillDetailItemPrice,
			OrderAmount: val.BillDetailItemAmount,
			CreatedAt:   val.CreatedAt,
		}
		best_sells = append(best_sells, best_sell)
	}
	for _, val := range bill {
		total_price += val.BillTotalPrice
		order_amount++
	}
	result := Schema.ReportHome{
		TotalPrice:  total_price,
		TotalProfit: total_profit,
		OrderAmount: order_amount,
		Summary:     total_price / order_amount,
		BestSell:    best_sells,
		Bill:        bill,
		BillDetail:  billDetail,
	}
	return &result, nil
}
