package repository

import (
	"coeffee/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type chatRepository struct {
	database   mongo.Database
	collection string
}


func NewChatRepository(database mongo.Database, collection string) domain.ChatRepository {
	return &chatRepository{
		database:   database,
		collection: collection}
}

// create chat

func (r *chatRepository) CreateChat(ctx context.Context , chat domain.Chat) (domain.Chat, error) {
	
	collection  := r.database.Collection(r.collection)

	_, err := collection.InsertOne(ctx, chat)

	if err != nil {

		return domain.Chat{}, err
}
	return chat, nil
}


// get all chat i have both sent and received id as parameter so fetch all chat that has either as sender or receiver
func (r *chatRepository) GetAllChat(ctx context.Context, senderID, receiverID string) ([]domain.Chat, error) {
	
	collection := r.database.Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.M{"$or": []bson.M{{"sender_id": senderID}, {"receiver_id": receiverID}}})

	if err != nil {
		return nil, err
	}

	var chats []domain.Chat

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var chat domain.Chat
		cursor.Decode(&chat)
		chats = append(chats, chat)
	}

	return chats, nil
}