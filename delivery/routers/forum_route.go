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

func NewForumRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	timeout := time.Duration(config.ContextTimeout) * time.Second

	repo := repository.NewForumRepository(DB, config.BlogCollection)
	replyrepo := repository.NewReplyRepository(DB, config.ReplyCollection)
	userrepo := repository.NewUserRepository(DB, config.UserCollection)

	usecase := usecase.NewBlogUseCase(repo, replyrepo , userrepo , timeout)

	forumController := controllers.NewForumController(usecase)

	


	forum := route.Group("/forum")
	forum.Use(infrastracture.AuthMiddleware())
	{
		forum.POST("/post", forumController.CreatePost)
		forum.GET("/getBlogById/:id", forumController.GetBlogByID)
		forum.GET("/getAllBlog", forumController.GetAllBlog)
		forum.POST("/:id/reply", forumController.CreateReply)
		forum.GET("/:id/getReply", forumController.GetReply)
		
	}
	resoure := route.Group("/resource")
	resoure.Use(infrastracture.AuthMiddleware())
	{
		resoure.POST("/post", forumController.CreateResource)
		resoure.GET("/getAllResource", forumController.GetAllResource)
		
	}



}