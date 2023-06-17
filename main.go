package main

import (
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

var client = initRedis()
var e = echo.New()

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
