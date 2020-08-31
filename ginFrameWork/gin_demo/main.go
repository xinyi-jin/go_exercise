package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/someGet", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get",
		})
	})

	router.POST("/somePost", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "POST",
		})
	})

	router.PUT("/somePut", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "PULL",
		})
	})

	router.DELETE("/delete", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, gin.H{
			"message": "DELETE",
		})
	})

	router.Run(":9090")
}
