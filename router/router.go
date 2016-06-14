package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/go_stock_with_gin/model"
	"fmt"
	"log"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("src/github.com/go_stock_with_gin/index.tmpl")
	r.GET("/", Index)
	r.GET("/build/:fileName", resource)
	r.GET("/testDB", testDB)

	InitQuoteRouter(r)
	InitUserRouter(r)
	InitialWatchListRouter(r)

	return r
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func resource(c *gin.Context){
	c.File("src/github.com/go_stock_with_gin/build/" + c.Param("fileName"))
}

func testDB(c *gin.Context){
	email := c.Query("email")
	user, err := model.GetUserByEmail(email)

	if err != nil{
		panic(err)
	}

	fmt.Println(user.Email)
	fmt.Println(user.Name)
	fmt.Println(user.Created)

	user.Name = "Tester2"

	if err := user.Save(); err != nil {
		log.Fatal(err)
	}

}


