package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("src/github.com/go_stock_with_gin/index.tmpl")
	r.GET("/", index)
	r.GET("/build/:fileName", resource)

	InitQuoteRouter(r)

	return r
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func resource(c *gin.Context){
	c.File("src/github.com/go_stock_with_gin/build/" + c.Param("fileName"))
}