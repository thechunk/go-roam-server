package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/garyburd/redigo/redis"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	redisHost := os.Getenv("REDIS_URL")
	if redisHost == "" {
		redisHost = "redis://127.0.0.1:6379"
	}
	c, err := redis.DialURL(redisHost)
	if (err == nil) {
		defer c.Close()
	}

	r.GET("/redis", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error": err,
		})
	})

	r.Run();
}
