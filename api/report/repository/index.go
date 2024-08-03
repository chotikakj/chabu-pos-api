package ReportRepository

import (
	"context"
	"pos-api/database"
	"pos-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetBillReport(start_date time.Time, end_date time.Time) ([]models.BillModel, error) {
	db := database.MongoDB
	ctx := context.Background()
	filter := bson.M{
		"created_at": bson.M{
			"$gte": start_date,
			"$lte": end_date,
		},
	}
	cursor, err := db.Collection("bill").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	bills := []models.BillModel{}
	for cursor.Next(ctx) {
		bill := models.BillModel{}
		err := cursor.Decode(&bill)
		if err != nil {
			return nil, err
		}
		bills = append(bills, bill)
	}
	return bills, nil
}

func GetBillDetailReport(start_date time.Time, end_date time.Time) ([]models.BillDetailModel, error) {
	db := database.MongoDB
	ctx := context.Background()
	filter := bson.M{
		"created_at": bson.M{
			"$gte": start_date,
			"$lte": end_date,
		},
	}
	cursor, err := db.Collection("bill-detail").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	bills := []models.BillDetailModel{}
	for cursor.Next(ctx) {
		bill := models.BillDetailModel{}
		err := cursor.Decode(&bill)
		if err != nil {
			return nil, err
		}
		bills = append(bills, bill)
	}
	return bills, nil
}
