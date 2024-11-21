package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// curl --location --globoff 'localhost:8080/post?ids[a]=1234&ids[b]=hello&ids[c]=111' \
// --form 'names[first]="111"' \
// --form 'names[second]="sss"'
func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v sssss", ids, names)
	})
	router.Run(":8080")
}
