package main

import "github.com/gin-gonic/gin"

func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件从任何 panic 中恢复，如果有，则写入 500。
	r.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	r.GET("/benchmark", ab, benchEndpoint)

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(AuthRequired())

	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

func analyticsEndpoint(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "analytics",
	})

}

func readEndpoint(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "read",
	})
}

func submitEndpoint(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "submit",
	})
}

func loginEndpoint(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "login",
	})
}

func AuthRequired() gin.HandlerFunc {
	return func(context *gin.Context) {
		// example logic
		// if token != "12345" {
		// 	context.AbortWithStatus(401)
		// 	return
		// }
	}
}

func ab(context *gin.Context) {
	context.Set("example", "12345")
}

func benchEndpoint(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "bench",
		"example": context.MustGet("example"),
	})
}
