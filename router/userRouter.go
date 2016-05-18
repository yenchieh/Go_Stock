package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go_stock_with_gin/controller"
)

func InitUserRouter(r *gin.Engine){
	r.POST("/user/create", controller.CreateUser)
	r.POST("/user/get/:email", controller.GetUserByEmail)
}

