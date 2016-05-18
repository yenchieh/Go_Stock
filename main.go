package main

import (
	"github.com/go_stock_with_gin/router"
	"github.com/go_stock_with_gin/common"
)

func main() {
	//Bootstrap Database
	common.InitDatabase()

	r := router.InitRoutes()
	r.Run(":8080")

}
