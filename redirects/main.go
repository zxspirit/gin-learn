package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// Issuing a HTTP redirect is easy. Both internal and external locations are supported.
	r.GET("/test", func(c *gin.Context) {
		//c.Redirect(http.StatusFound, "http://www.google.com/")
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})
	// Issuing a HTTP redirect from POST. Refer to issue: #444
	r.POST("/test", func(c *gin.Context) {
		//c.Redirect(http.StatusMovedPermanently, "/foo")
		c.Redirect(http.StatusFound, "/foo")
	})
	r.GET("/foo", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})
	// Issuing a Router redirect, use HandleContext like below.
	r.GET("/test1", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
	fmt.Println("Server is running on port 8080")
}
