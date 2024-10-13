package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Message    string `json:"message"`
	CreatedAt  string `json:"created_at"`
}

type ChatRepository interface {
	CreateChat(ctx context.Context, chat Chat) (Chat, error)
	GetAllChat(ctx context.Context, senderID, receiverID string) ([]Chat, error)
	// GetAllChat(senderID, receiverID string) ([]Chat, error)
}

type ChatUseCase interface {
	CreateChat(chat Chat) (Chat, ErrorResponse)
	GetAllChat(senderID, receiverID string) ([]Chat, error)
}
