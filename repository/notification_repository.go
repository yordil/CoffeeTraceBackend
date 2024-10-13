package repository

import (
	"coeffee/domain"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type notificationRepository struct {
	database   mongo.Database
	collection string
}


func NewNotificationRepository(database mongo.Database, collection string) domain.NotificationRepository {
	return &notificationRepository{
		database:   database,
		collection: collection}
}


func (r *notificationRepository) CreateNotification(ctx context.Context , notification domain.Notification) error {
	_, err := r.database.Collection(r.collection).InsertOne(ctx, notification)
	if err != nil {
		return err
	}
	return nil


}

// UpdateNotification

func (r *notificationRepository) UpdateNotification(ctx context.Context , notification domain.Notification) error {

	return nil

}