package main

import (
	"os"

	"fmt"

	"github.com/urfave/cli"
	"github.com/yenchieh/Go_Stock/config"
	"github.com/yenchieh/Go_Stock/controller"
	"github.com/yenchieh/Go_Stock/router"
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
			Value:  "go",
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
		cli.StringFlag{
			EnvVar: "ALPHA_VANTAGE_KEY",
			Name:   "alpha_vantage_key",
			Usage:  "For retrieve data from alpha vantage",
		},
	}

	app.Action = func(c *cli.Context) error {
		controller.SetupDatabase(controller.Database{
			Name:     c.String("database_name"),
			User:     c.String("database_user"),
			Password: c.String("database_password"),
			IP:       c.String("database_IP"),
		})

		config.SetupEnv(config.Environment{
			Debug:           c.Bool("debug"),
			AlphaVantageKey: c.String("alpha_vantage_key"),
		})

		db := controller.NewDB()
		defer db.Close()

		r := router.New()

		return r.Run(fmt.Sprintf(":%s", c.String("port")))
	}

	app.Run(os.Args)

}
