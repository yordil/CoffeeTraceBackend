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


func NewUserRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database){

	timeout := time.Duration(config.ContextTimeout) * time.Second
	
	repo := repository.NewUserRepository(DB, config.UserCollection)


	tokenGen := infrastracture.NewTokenGenerator()
	passwordSvc := infrastracture.NewPasswordService()


	usecase := usecase.NewUserUseCase(repo , timeout, tokenGen, passwordSvc)


	


	userController := controllers.NewUserController(usecase)

	user := route.Group("/user")
	{
		user.POST("/register", userController.CreateAccount)
		user.POST("/login", userController.Login)
		user.Use(infrastracture.AuthMiddleware())
		user.POST("/ID", userController.GetByID)
		user.POST("/updateProfile", userController.UpdateProfile)
		user.POST("/add-driver" , userController.AddDriver)
		
		user.GET("/get-all", userController.GetAllUser)
		user.GET("/me" , userController.GetMe)
		
	}

}