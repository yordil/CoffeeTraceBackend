package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Name             string    `json:"farmer_id"`
	Type             string    `json:"type"`
	Quantity         float64    `json:"quantity"`
	ProductID 		string 		`json:"product_id"`
	Shipping  	     string   `json:"shipping"`
	Status 		    string    `json:"status"`
	MerchatName 	 string `json:"merchant_name"`
	MerchatAdress 	 string `json:"merchant_address"`
	FarmerName 		 string `json:"farmer_name"`
	FarmerAdress 	 string `json:"farmer_address"`
	NotificationDate time.Time `json:"notification_date"`
	
}



type NotificationRepository interface {
	CreateNotification(ctx context.Context , notification Notification) error
	UpdateNotification(ctx context.Context , notification Notification) error
}

