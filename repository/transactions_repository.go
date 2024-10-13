package repository

import (
	"coeffee/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type transactionsRepository struct {
	database   mongo.Database
	collection string
}

func NewTransactionRepository(database mongo.Database, collection string) domain.TransactionsRepository {
	return &transactionsRepository{
		database:   database,
		collection: collection}

}

func (r *transactionsRepository) CreateTransaction(transaction domain.Transactions) (domain.Transactions, error) {
	transaction.ID = primitive.NewObjectID()
	_, err := r.database.Collection(r.collection).InsertOne(context.Background(), transaction)
	if err != nil {
		return domain.Transactions{}, err
	}

	return transaction, nil
}

func (r *transactionsRepository) GetTransactions(userID, role string) ([]domain.Transactions, error) {
	var transactions []domain.Transactions
	filter := bson.M{}
	if role == "farmer" {
		filter = bson.M{"farmerid": userID}
	} else {
		filter = bson.M{"merchantid": userID}
	}

	cursor, err := r.database.Collection(r.collection).Find(context.Background(), filter)
	if err != nil {
		return []domain.Transactions{}, err
	}

	if err = cursor.All(context.Background(), &transactions); err != nil {
		return []domain.Transactions{}, err
	}

	return transactions, nil
}