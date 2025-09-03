package main

import (
	"test-jwt-go/config"
	"test-jwt-go/controllers"
	"test-jwt-go/database"
	"test-jwt-go/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadAppConfig()

	database.Connect(config.AppConfig.ConnectionString, config.AppConfig.Config...)
	database.Migrate()

	router := initRouter()
	router.Run(":8080")

}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	return router
}
