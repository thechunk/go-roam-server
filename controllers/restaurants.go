package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thechunk/roam-server/database"
	"strconv"
)

func RestaurantsNearbyController() func(c *gin.Context) {
	return func(c *gin.Context) {
		lat, err := strconv.ParseFloat(c.Query("lat"), 64)
		lng, err := strconv.ParseFloat(c.Query("lng"), 64)
		rad, err := strconv.ParseFloat(c.Query("rad"), 64)
		if err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		if rest, err := database.RestaurantsNearby(lat, lng, rad); err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"data": rest,
			})
		}
	}
}

func RestaurantByIdController() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		if rest, err := database.RestaurantById(id); err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"data": rest,
			})
		}
	}
}
