package mappings

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine;

func CreateUrlMappings() {
	Router = gin.Default()
	api := Router.Group("/api")
	{
		api.GET("/test",func(ctx *gin.Context){
			ctx.JSON(200, gin.H{
			 "message": "test successful",
			});
		})
	}
}
