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
		redisHost = ":6379"
	}
	c, err := redis.Dial("tcp", redisHost)
	_ = c

	r.GET("/redis", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error": err,
		})
	})

	r.Run();
}
