package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BillModel struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	BillTotalPrice int                `bson:"bill_total_price" json:"bill_total_price"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
}
