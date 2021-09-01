package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)

	router.Run()
}

func posting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
		"method":  "POST",
	})
}

func getting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
		"method":  "GET",
	})
}
