package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
		v1.GET("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.GET("/login", loginEndpoint)
		v2.GET("/submit", submitEndpoint)
		v2.GET("/read", readEndpoint)
	}

	router.Run(":8080")
}

func readEndpoint(context *gin.Context) {
	context.JSON(200, gin.H{"message": "read"})
}

func submitEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{"message": "submit"})
}

func loginEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{"message": "login"})
}
