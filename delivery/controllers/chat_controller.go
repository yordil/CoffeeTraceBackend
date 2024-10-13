package controllers

import (
	"coeffee/domain"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ChatController struct {
	ChatUsecase domain.ChatUseCase
}

func NewChatController(bc domain.ChatUseCase) *ChatController {
	return &ChatController{
		ChatUsecase: bc,
	}
}

// CreateChat

func (cc *ChatController) CreateChat(c *gin.Context) {
	var chat domain.Chat

	// get the id of the current logged in user from the context
	userID := c.GetString("user_id")
	
	
	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	chat.SenderID = userID
	chat.ReceiverID = c.Param("id")

	suss, err := cc.ChatUsecase.CreateChat(chat)

	if err.Message != "" {
		c.JSON(err.Status, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "created chat successfully" , "chat": suss})



}

// get all chat i have both sent and received id as parameter so fetch all chat that has either as sender or receiver
// GetAllChat
func (cc *ChatController) GetAllChat(c *gin.Context) {
	senderID := c.GetString("user_id")
	receiverID := c.Param("id")

	chats, err := cc.ChatUsecase.GetAllChat(senderID, receiverID)

	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to get chat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"chats": chats})

}