package common

import (
	"fmt"
	"github.com/go_stock_with_gin/data"
	"database/sql"
	"log"
)

func InitDatabase(){
	db := data.GetDatabase()
	initUserTable(db)
	initStockTable(db)

	defer db.Close()

}

func initUserTable(db *sql.DB){
	/* Table creation */
	// If user table is not exist, create new User Table
	if _, err := db.Query("SELECT * FROM user"); err != nil {
		createUserTableQuery := "CREATE TABLE user (id INT(10) NOT NULL AUTO_INCREMENT, name VARCHAR(64) NOT NULL, email VARCHAR(200) NOT NULL, password VARCHAR(255), created TIMESTAMP NULL DEFAULT NULL, PRIMARY KEY (id))"
		fmt.Printf("Creating User Table: \n%s\n", createUserTableQuery)

		if _, err := db.Exec(createUserTableQuery);  err != nil{
			panic(err.Error())
		}
		fmt.Printf("-- User Table Created --\n")
	}

	// If user role table is not exist, create new User Role Table
	if _, err := db.Query("SELECT * FROM user_role"); err != nil {
		createUserRoleQuery := "CREATE TABLE user_role (id INT(10) NOT NULL AUTO_INCREMENT, authority VARCHAR(20) NOT NULL, PRIMARY KEY (id))"
		fmt.Printf("Creating User Role: \n%s\n", createUserRoleQuery)
		if _, err := db.Exec(createUserRoleQuery);  err != nil{
			panic(err.Error())
		}

		fmt.Printf("-- User Role Table Created --\n")
	}

	// If user role connection table is not exist, create new User Role connection table
	if _, err := db.Query("SELECT * FROM user_role_connection"); err != nil {
		createUserRoleQuery := "CREATE TABLE user_role_connection (id INT(10) NOT NULL AUTO_INCREMENT, user_id INT(10) NOT NULL, role_id INT(10) NOT NULL, PRIMARY KEY (id), FOREIGN KEY (user_id) REFERENCES user(id), FOREIGN KEY (role_id) REFERENCES user_role(id))"

		fmt.Printf("Creating User Role connection: \n%s\n", createUserRoleQuery)
		if _, err := db.Exec(createUserRoleQuery);  err != nil{
			panic(err.Error())
		}

		fmt.Printf("-- User Role Connection Table Created --\n")
	}
	/* Row Creation */

	rows := db.QueryRow("SELECT * FROM user_role WHERE authority = ?", "ROLE_USER")
	err := rows.Scan()

	if err == sql.ErrNoRows {
		stmt, err := db.Prepare("INSERT user_role SET authority = ?")

		if err != nil {
			log.Fatal(err)
		}

		if _, err = stmt.Exec("ROLE_USER"); err != nil {
			log. Fatal(err)
		}
	}

}

func initStockTable(db *sql.DB){

	// If user WatchList table is not exist, create WatchList table
	if _, err := db.Query("SELECT * FROM stocks"); err != nil {
		createWatchListQuery := "CREATE TABLE stocks (id INT(10) NOT NULL AUTO_INCREMENT, name VARCHAR(100), symbol VARCHAR(10) NOT NULL, PRIMARY KEY (id))"

		fmt.Printf("Creating stock_list: \n%s\n", createWatchListQuery)
		if _, err := db.Exec(createWatchListQuery);  err != nil{
			panic(err.Error())
		}

		fmt.Printf("-- User stock_list Table Created --\n")
	}

	// If user Watch List connection table is not exist, create WatchList connection table
	if _, err := db.Query("SELECT * FROM user_stocks_connection"); err != nil {
		createWatchListConnectionQuery := "CREATE TABLE user_stocks_connection (id INT(10) NOT NULL AUTO_INCREMENT, user_id INT(10) NOT NULL, stock_id INT(10) NOT NULL, type VARCHAR(10) NOT NULL, date_created TIMESTAMP NOT NULL, PRIMARY KEY (id), FOREIGN KEY (user_id) REFERENCES user(id), FOREIGN KEY (stock_id) REFERENCES stocks(id))"

		fmt.Printf("Creating user_stocks_connection: \n%s\n", createWatchListConnectionQuery)
		if _, err := db.Exec(createWatchListConnectionQuery);  err != nil{
			panic(err.Error())
		}

		fmt.Printf("-- User user_stocks_connection Table Created --\n")
	}
}