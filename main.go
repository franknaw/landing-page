package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//err-os := os.Setenv("T_PATH", "/myapp")
	//if err-os != nil {
	//	return
	//}

	router := gin.Default()
	//router.LoadHTMLGlob(filepath.Join(os.Getenv("T_PATH"), "templates/*"))
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Landing Page",
		})
	})
	err := router.Run(":8080")
	if err != nil {
		return
	}

}
