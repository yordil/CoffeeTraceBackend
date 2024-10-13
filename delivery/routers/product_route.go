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

func NewProductRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	timeout := time.Duration(config.ContextTimeout) * time.Second

	repo := repository.NewProductRepository(DB, config.ProductCollection)
	usecase := usecase.NewProductUseCase(repo, timeout)

	productHandler := controllers.ProductController{
		ProductUsecase: usecase,
	}


	product := route.Group("/product")
	product.Use(infrastracture.AuthMiddleware())
	{
		product.POST("/create", productHandler.CreateProduct)
		product.GET("/getall", productHandler.GetAllProduct)
	}
}