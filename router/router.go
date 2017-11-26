package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yenchieh/Go_Stock/controller"
)

func New() *gin.Engine {
	r := gin.Default()

	r.Static("/assets", "view/dist/")
	r.LoadHTMLGlob("view/dist/*.html")
	r.GET("/", index)

	// APIs
	r.GET("/searchStock", controller.SearchStock)

	return r
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
