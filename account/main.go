package main

import (
	"log"

	"github.com/On-A-Rocket/Authorization-System/account/application/command"
	"github.com/On-A-Rocket/Authorization-System/account/config"
	"github.com/On-A-Rocket/Authorization-System/account/controller"
	"github.com/On-A-Rocket/Authorization-System/account/domain/entity"
	"github.com/On-A-Rocket/Authorization-System/account/repository"
	"github.com/gin-gonic/gin"
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

func main() {
	config := config.Initialize()

	router := gin.Default()

	dbConnection := getDatabaseConnection(config)

	repository := repository.NewRepository(dbConnection)
	command := command.NewCommand(repository)

	ctl := controller.NewController(*command)
	ctl.Routing(router)

	log.Fatal(router.Run(":5001"))
}
