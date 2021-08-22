package main

import (
	"log"

	"github.com/On-A-Rocket/Authorization-System/auth/application/command"
	"github.com/On-A-Rocket/Authorization-System/auth/application/query"
	"github.com/On-A-Rocket/Authorization-System/auth/config"
	"github.com/On-A-Rocket/Authorization-System/auth/controller"
	"github.com/On-A-Rocket/Authorization-System/auth/docs"
	"github.com/On-A-Rocket/Authorization-System/auth/repository"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func setSwaggerInfo() {
	docs.SwaggerInfo.Title = "Authorization System REST API"
	docs.SwaggerInfo.Description = "This is a Authorization System rest api server for swagger"
	docs.SwaggerInfo.Version = "1.0"
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(corsMiddleware())

	setSwaggerInfo()
	config := config.Initialize()

	dbConnection := config.Database().Connection()
	_ = config.Redis().Client()

	query := query.NewQuery(dbConnection, config)

	repository := repository.NewRepository(dbConnection)
	command := command.NewCommand(repository, config)

	ctl := controller.NewController(*query, *command, config)
	ctl.Routing(router)

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(router.Run(":5001"))
}
