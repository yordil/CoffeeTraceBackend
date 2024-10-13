package routers

import (
	"coeffee/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	NewUserRouter(route , config, DB)
	NewProductRouter(route , config, DB)
	NewOrderRouter(route , config, DB)
	NewTransactionRouter(route , config, DB)
	NewForumRouter(route , config, DB)
}