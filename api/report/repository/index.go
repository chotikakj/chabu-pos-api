package ReportRepository

import (
	"context"
	"pos-api/database"
	"pos-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	cursor, err := db.Collection("bill").Find(ctx, filter, options.Find().SetSort(bson.D{{"created_at", -1}}))
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

func GetBillDetailByBillID(BillID string) ([]models.BillDetailModel, error) {
	db := database.MongoDB
	ctx := context.Background()
	bill_id, _ := primitive.ObjectIDFromHex(BillID)
	cursor, err := db.Collection("bill-detail").Find(ctx, bson.M{"bill_id": bill_id})
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
