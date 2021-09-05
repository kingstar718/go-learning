package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}

func readEndpoint(context *gin.Context) {

}

func submitEndpoint(context *gin.Context) {

}

func loginEndpoint(context *gin.Context) {

}
