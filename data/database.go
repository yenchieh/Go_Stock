package data

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)


// connect to our database server with data source name
// data source name configuration has the following parameters :
// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]

// example config :
// user:password@tcp(127.0.0.1:3306)/database
func GetDatabase() *sql.DB{
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_stock?charset=utf8&parseTime=true")
	if err != nil{
		log.Fatal(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	return db
}
