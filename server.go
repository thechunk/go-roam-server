package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thechunk/roam-server/controllers"
)

type Engine struct {
}

func (r *Engine) GET(relativePath string, cb func(c *gin.Context)) []interface{} {
	return nil
}

type EngineInterface interface {
	GET(relativePath string, cb ...gin.HandlerFunc) gin.IRoutes
	Run(s ...string) error
}

func startServer(r EngineInterface) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/v1/restaurants", controllers.RestaurantsNearbyController())
	r.GET("/api/v1/restaurants/:id", controllers.RestaurantByIdController())

	r.Run()
}

func main() {
	r := gin.Default()
	startServer(r)
}
