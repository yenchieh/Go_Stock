package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/dist", "view/dist")
	r.LoadHTMLGlob("view/dist/*.html")
	r.GET("/", index)

	r.Run(":14443")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
