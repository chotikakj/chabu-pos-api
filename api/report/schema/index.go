package ReportSchema

import (
	"pos-api/models"
	"time"
)

type ReportHomeDto struct {
	StartDate time.Time `json:"start_date" validate:"required"`
	EndDate   time.Time `json:"end_date" validate:"required"`
}

type ReportHome struct {
	TotalPrice  int                      `json:"total_price"`
	TotalProfit int                      `json:"total_profit"`
	OrderAmount int                      `json:"order_amount"`
	Summary     int                      `json:"summary"`
	BestSell    []BestSellItems          `json:"best_sell"`
	Bill        []models.BillModel       `json:"bill"`
	BillDetail  []models.BillDetailModel `json:"bill_detail"`
}

type BestSellItems struct {
	Name        string    `json:"name"`
	Pricing     int       `json:"pricing"`
	OrderAmount int       `json:"order_amount"`
	CreatedAt   time.Time `json:"created_at"`
}
