package routers

import (
	"coeffee/config"
	"coeffee/delivery/controllers"
	"coeffee/infrastracture"
	"coeffee/repository"
	"coeffee/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTransactionRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {
	timeout := time.Duration(config.ContextTimeout) * time.Second
	repo := repository.NewTransactionRepository(DB, config.TransactionCollection)
	usecase := usecase.NewTransactionsUseCase(repo, timeout)
	transactionController := controllers.NewTransactionController(usecase)
	transaction := route.Group("/transaction")
	transaction.Use(infrastracture.AuthMiddleware())
	{
		transaction.POST("/create", transactionController.CreateTransaction)
		transaction.GET("/getmytransactions", transactionController.GetMyTransactions)
		
	}
}