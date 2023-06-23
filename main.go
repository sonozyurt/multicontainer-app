package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var client = initRedis()
var e = echo.New()
var db = connPostgres()

func main() {
	e.GET("/", mainPage)
	e.POST("/", postValue)
	e.Start(":8080")

}

func initRedis() *redis.Client {
	client := redis.NewClient(
		&redis.Options{
			Addr:       "127.0.0.1:6379",
			Password:   "",
			DB:         0,
			MaxRetries: 1000,
		})
	return client
}

func connPostgres() *gorm.DB {
	dsn := "host=0.0.0.0 port=5432 user=postgres password=postgres dbname=database sslmode=disable"
	//urlExample := "postgres://postgres:postgres@localhost:5432/database"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = conn.AutoMigrate(&dbData{})
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
