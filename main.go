package main

import (
	"net/http"
	"os"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"github.com/yenchieh/Go_Stock/controller"
)

func main() {
	app := cli.NewApp()
	app.Name = "Go Stock"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			EnvVar: "DEBUG",
			Name:   "debug",
			Usage:  "Debug",
			Value:  "false",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_USER",
			Name:   "database_user",
			Usage:  "Database user name",
			Value:  "yenchieh",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_PASSWORD",
			Name:   "database_password",
			Usage:  "Database password",
			Value:  "aaaaaa",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_IP",
			Name:   "database_IP",
			Usage:  "Database IP address with port number",
			Value:  "127.0.0.1",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_NAME",
			Name:   "database_name",
			Usage:  "Database name",
			Value:  "go_stock",
		},
		cli.StringFlag{
			EnvVar: "PORT",
			Name:   "port",
			Value:  "8120",
			Usage:  "",
		},
	}

	app.Action = func(c *cli.Context) error {
		controller.SetupDatabase(controller.Database{
			Name:     c.String("database_name"),
			User:     c.String("database_user"),
			Password: c.String("database_password"),
			IP:       c.String("database_IP"),
		})

		db := controller.NewDB()
		defer db.Close()

		r := gin.Default()

		r.Static("/dist", "view/dist")
		r.LoadHTMLGlob("view/dist/*.html")
		r.GET("/", index)

		return r.Run(fmt.Sprintf(":%s", c.String("port")))
	}

	app.Run(os.Args)

}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
