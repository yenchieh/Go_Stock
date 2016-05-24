package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go_stock_with_gin/controller"
)

func InitialWatchListRouter(r *gin.Engine){
	r.POST("/watchList/add", controller.AddStockToWatchList)
}
