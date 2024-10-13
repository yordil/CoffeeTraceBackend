package routers

import (
	"coeffee/config"
	"coeffee/delivery/controllers"
	"coeffee/infrastracture"
	"coeffee/repository"
	"coeffee/usecase"
	"time"

	// "coefee/infrastructure"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewChatRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	timeout := time.Duration(config.ContextTimeout) * time.Second

	chatRepo := repository.NewChatRepository(DB, "chat")

	chatUseCase := usecase.NewChatUseCase(chatRepo, timeout)

	chatHandler := controllers.NewChatController(chatUseCase)
	chat := route.Group("/chat")
	chat.Use(infrastracture.AuthMiddleware())

	{
		chat.POST("/", chatHandler.CreateChat)
	// get all chat
		chat.GET("/:id", chatHandler.GetAllChat)
	}


	


}