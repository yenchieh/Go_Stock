package main

import (
	"github.com/go_stock_with_gin/router"
)

func main() {

	r := router.InitRoutes()
	r.Run(":8080")

}
