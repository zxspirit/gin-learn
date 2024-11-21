package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

//	func main() {
//		router := gin.Default()
//		router.LoadHTMLGlob("templates/*")
//		//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
//		router.GET("/index", func(c *gin.Context) {
//			c.HTML(http.StatusOK, "index.tmpl", gin.H{
//				"title": "Main website",
//			})
//		})
//		router.Run(":8080")
//	}

func main() {
	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
			"now":   time.Date(2017, 0o7, 0o1, 0, 0, 0, 0, time.UTC),
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	router.Run(":8080")
}
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

//func main() {
//	router := gin.Default()
//	html := template.Must(template.ParseFiles("templates/posts/index.tmpl", "templates/users/index.tmpl"))
//	router.SetHTMLTemplate(html)
//	router.GET("/posts/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
//			"title": "Posts",
//		})
//	})
//	router.GET("/users/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
//			"title": "Users",
//		})
//	})
//	router.Run(":8080")
//}
