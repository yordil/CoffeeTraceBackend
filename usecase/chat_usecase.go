package usecase

import (
	"coeffee/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatUseCase struct {
	ChatRepository		domain.ChatRepository
	contextTimeout  	time.Duration
}

func NewChatUseCase(chat domain.ChatRepository,  timeout time.Duration ) domain.ChatUseCase {
	return &ChatUseCase{
		ChatRepository: chat,
		contextTimeout: timeout,
	}
}

func (uc *ChatUseCase) CreateChat(chat domain.Chat) (domain.Chat, domain.ErrorResponse) {
	
	ctx , cancel := context.WithTimeout(context.Background(), uc.contextTimeout)

	defer cancel()	

	if chat.SenderID == "" || chat.ReceiverID == "" || chat.Message == "" {
		return domain.Chat{}, domain.ErrorResponse{Message: "SenderID, ReceiverID or Message cannot be empty", Status: 400}
	}

	chat.ID = primitive.NewObjectID()
	_, err := uc.ChatRepository.CreateChat(ctx, chat)
	chat.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	if err != nil {
		return domain.Chat{}, domain.ErrorResponse{Message: "Failed to create chat", Status: 500}
	}

	return chat, domain.ErrorResponse{}
}

func (uc *ChatUseCase) GetAllChat(senderID, receiverID string) ([]domain.Chat, error) {
	ctx , cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()

	chats, err := uc.ChatRepository.GetAllChat(ctx, senderID, receiverID)

	if err != nil {
		return nil, err
	}

	return chats, nil
	
}