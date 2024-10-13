package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transactions struct {
	ID           primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Amount       int                `json:"amount"`
	Description  string             `json:"description"`
	OrderID      string                `json:"order_id"`
	FarmerID     string             `json:"farmer_id"`
	FarmerName   string             `json:"farmer_name"`
	MerchantID   string             `json:"merchant_id"`
	MerchantName string             `json:"merchant_name"`
	Quantity     int                `json:"quantity"`
	TotalPrice   int                `json:"total_price"`
	Datetime     string             `json:"datetime"`
}

type TransactionsRepository interface {
	CreateTransaction(transaction Transactions) (Transactions, error)
	GetTransactions(userID, role string) ([]Transactions, error)
}

type TransactionsUseCase interface {
	CreateTransaction(transaction Transactions) (Transactions, error)
	GetTransactions(userId, role string) ([]Transactions, error)
}