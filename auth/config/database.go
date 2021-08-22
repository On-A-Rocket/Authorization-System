package config

import (
	"log"
	"os"

	"github.com/On-A-Rocket/Authorization-System/auth/domain/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseInterface interface {
	Port() string
	Host() string
	Name() string
	User() string
	Password() string
	Connection() *gorm.DB
}

type Database struct {
	port     string
	host     string
	name     string
	user     string
	password string
}

func newDatabaseConfig() *Database {
	db := &Database{
		port:     "3306",
		host:     "localhost",
		name:     "test",
		user:     "root",
		password: "1",
	}

	if port := os.Getenv("DATABASE_PORT"); port != "" {
		db.port = port
	}
	if host := os.Getenv("DATABASE_HOST"); host != "" {
		db.host = host
	}
	if name := os.Getenv("MYSQL_DATABASE"); name != "" {
		db.name = name
	}
	if user := os.Getenv("MYSQL_USER"); user != "" {
		db.user = user
	}
	if password := os.Getenv("MYSQL_PASSWORD"); password != "" {
		db.password = password
	}

	return db
}

func (db *Database) Connection() *gorm.DB {
	user := db.User()
	password := db.Password()
	host := db.Host()
	port := db.Port()
	name := db.Name()
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	dbConnection.AutoMigrate(
		&entity.Account{},
	)

	return dbConnection
}

func (db *Database) Port() string {
	return db.port
}

func (db *Database) Host() string {
	return db.host
}

func (db *Database) Name() string {
	return db.name
}

func (db *Database) User() string {
	return db.user
}

func (db *Database) Password() string {
	return db.password
}
