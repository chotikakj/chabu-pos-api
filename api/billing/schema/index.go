package BillingSchema

type CreateBillDto struct {
	BillTotalPrice int              `json:"bill_total_price" validate:"required"`
	BillDetail     []BillDetailData `json:"bill_detail" validate:"required"`
}

type BillDetailData struct {
	BillDetailItemName   string `json:"bill_detail_item_name" validate:"required"`
	BillDetailItemType   string `json:"bill_detail_item_type" validate:"required"`
	BillDetailItemPrice  int    `json:"bill_detail_item_price" validate:"required"`
	BillDetailItemAmount int    `json:"bill_detail_item_amount" validate:"required"`
	BillDetailItemCost   int    `json:"bill_detail_item_cost" validate:"required"`
}
