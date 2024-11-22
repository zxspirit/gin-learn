package main

import "github.com/gin-gonic/gin"

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/some", getting)
	router.POST("/some", posting)
	router.PUT("/some", putting)
	router.DELETE("/some", deleting)
	router.PATCH("/some", patching)
	router.HEAD("/some", head)
	router.OPTIONS("/some", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}

func options(context *gin.Context) {
	context.String(200, "OPTIONS request")
}

func head(context *gin.Context) {
	context.String(200, "HEAD request")
}

func patching(context *gin.Context) {
	context.String(200, "PATCH request")
}

func deleting(context *gin.Context) {
	context.String(200, "DELETE request")
}

func putting(context *gin.Context) {
	context.String(200, "PUT request")
}

func posting(context *gin.Context) {
	context.String(200, "POST request")
}

func getting(context *gin.Context) {
	context.String(200, "GET request")
}
