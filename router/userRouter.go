package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go_stock_with_gin/controller"
)

func InitUserRouter(r *gin.Engine){
	r.POST("/user/registerOrLogin", controller.RegisterOrLogin)
	//r.POST("/user/create", controller.CreateUser)
	r.GET("/user/get/:email", controller.GetUserByEmail)
	r.GET("/user/check/:email", controller.CheckEmail)
}

