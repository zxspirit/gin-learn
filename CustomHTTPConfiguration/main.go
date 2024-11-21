package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//func main() {
//	router := gin.Default()
//	router.GET("/", a)
//	http.ListenAndServe(":8080", router)
//}

func a(context *gin.Context) {
	context.String(http.StatusOK, "Hello World")
}

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()

	router.GET("/", a)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
