package BillingRepository

import (
	"context"
	Schema "pos-api/api/billing/schema"
	"pos-api/database"
	"pos-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBill(bill_total_price int) (primitive.ObjectID, error) {
	db := database.MongoDB
	ctx := context.Background()
	result, err := db.Collection("bill").InsertOne(ctx, models.BillModel{
		BillTotalPrice: bill_total_price,
		CreatedAt:      time.Now(),
	})
	if err != nil {
		return primitive.NilObjectID, err
	}
	ID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, err
	}
	return ID, nil
}

func CreateBillDetail(bill_detail []Schema.BillDetailData, billID primitive.ObjectID) error {
	db := database.MongoDB
	ctx := context.Background()
	var billInterfaces []interface{}
	for _, bill := range bill_detail {
		billInterfaces = append(billInterfaces, bson.D{
			{Key: "bill_detail_item_name", Value: bill.BillDetailItemName},
			{Key: "bill_detail_item_type", Value: bill.BillDetailItemType},
			{Key: "bill_detail_item_price", Value: bill.BillDetailItemPrice},
			{Key: "bill_detail_item_amount", Value: bill.BillDetailItemAmount},
			{Key: "bill_detail_item_total_price", Value: bill.BillDetailItemAmount * bill.BillDetailItemPrice},
			{Key: "bill_detail_item_cost", Value: bill.BillDetailItemCost},
			{Key: "bill_detail_item_total_cost", Value: bill.BillDetailItemAmount * bill.BillDetailItemCost},
			{Key: "created_at", Value: time.Now()},
			{Key: "updated_at", Value: time.Now()},
			{Key: "bill_id", Value: billID},
		})
	}
	_, err := db.Collection("bill-detail").InsertMany(ctx, billInterfaces)
	if err != nil {
		return err
	}
	return nil
}
