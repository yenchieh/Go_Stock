package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go_stock_with_gin/controller"
)

func InitialWatchListRouter(r *gin.Engine){
	r.GET("/watchList", Index)
	r.POST("/watchList/add", controller.AddStockToWatchList)
	r.GET("/watchList/getByEmail", controller.UserWatchListByEmail)
}
