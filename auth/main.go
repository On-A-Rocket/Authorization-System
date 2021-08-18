package main

import (
	"log"

	"github.com/On-A-Rocket/Authorization-System/auth/application/command"
	"github.com/On-A-Rocket/Authorization-System/auth/config"
	"github.com/On-A-Rocket/Authorization-System/auth/controller"
	"github.com/On-A-Rocket/Authorization-System/auth/docs"
	"github.com/On-A-Rocket/Authorization-System/auth/domain/entity"
	"github.com/On-A-Rocket/Authorization-System/auth/repository"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDatabaseConnection(config config.Interface) *gorm.DB {
	user := config.Database().User()
	password := config.Database().Password()
	host := config.Database().Host()
	port := config.Database().Port()
	name := config.Database().Name()
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&entity.Account{},
	)

	return db
}

func setSwaggerInfo() {
	docs.SwaggerInfo.Title = "Authorization System REST API"
	docs.SwaggerInfo.Description = "This is a Authorization System rest api server for swagger"
	docs.SwaggerInfo.Version = "1.0"
}

func main() {
	setSwaggerInfo()
	config := config.Initialize()

	router := gin.Default()

	dbConnection := getDatabaseConnection(config)

	repository := repository.NewRepository(dbConnection)
	command := command.NewCommand(repository)

	ctl := controller.NewController(*command)
	ctl.Routing(router)

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(router.Run(":5001"))
}
