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

func NewOrderRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	timeout := time.Duration(config.ContextTimeout) * time.Second


	transactionrepo := repository.NewTransactionRepository(DB, config.TransactionCollection)





	repo := repository.NewOrderRepository(DB, config.OrderCollection)
	Notifyrepo := repository.NewNotificationRepository(DB, config.NotificationCollection)
	Userrepo := repository.NewUserRepository(DB, config.UserCollection)
	prodrepo := repository.NewProductRepository(DB, config.ProductCollection)
	usecase := usecase.NewOrderUseCase(repo,Notifyrepo , Userrepo ,prodrepo, transactionrepo,  timeout)

	orderController := controllers.NewOrderController(usecase)


	order := route.Group("/order")
	order.Use(infrastracture.AuthMiddleware())
	{
		order.POST("/create", orderController.CreateOrder)
		order.GET("/getmyorders", orderController.GetMyOrders)


		// farmer

		order.POST("/farmer/:id/accept", orderController.AcceptOrderFarmer)
		order.POST("/farmer/:id/reject", orderController.RejectOrderFarmer)


		// driver
		order.POST("/driver/:id/accept", orderController.AcceptOrderDriver)
		order.POST("/driver/:id/reject", orderController.RejectOrderDriver)
		order.POST("/driver/:id/pickup", orderController.PickupOrderDriver)
		order.POST("/driver/:id/destinations", orderController.UpdateDestinations)


		// buyer
		order.POST("/buyer/:id/payFarmer", orderController.PayFarmer)
		order.POST("/buyer/:id/paydriver", orderController.PayDriver)

	}

}
