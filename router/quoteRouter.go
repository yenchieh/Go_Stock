package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go_stock_with_gin/controller"
)

func InitQuoteRouter(r *gin.Engine){
	r.GET("/getQuoteBySymbol", controller.GetQuoteBySymbol)

}