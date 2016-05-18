package common

import (
	"fmt"
	"github.com/go_stock_with_gin/data"
	"database/sql"
	"log"
)

func InitDatabase(){
	db := data.GetDatabase()
	defer db.Close()

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