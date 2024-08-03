package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BillDetailModel struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	BillDetailItemName       string             `bson:"bill_detail_item_name" json:"bill_detail_item_name"`
	BillDetailItemType       string             `bson:"bill_detail_item_type" json:"bill_detail_item_type"`
	BillDetailItemPrice      int                `bson:"bill_detail_item_price" json:"bill_detail_item_price"`
	BillDetailItemAmount     int                `bson:"bill_detail_item_amount" json:"bill_detail_item_amount"`
	BillDetailItemTotalPrice int                `bson:"bill_detail_item_total_price" json:"bill_detail_item_total_price"`
	BillDetailItemCost       int                `bson:"bill_detail_item_cost" json:"bill_detail_item_cost"`
	BillDetailItemTotalCost  int                `bson:"bill_detail_item_total_cost" json:"bill_detail_item_total_cost"`
	CreatedAt                time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt                time.Time          `bson:"updated_at" json:"updated_at"`
	BillID                   primitive.ObjectID `bson:"bill_id" json:"bill_id"`
}
