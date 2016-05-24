package model

import (
	"github.com/go_stock_with_gin/data"
	"fmt"
	"time"
	"crypto/sha256"
	"log"
	"database/sql"
)

type (
	User struct {
		Id       uint32 `json:"_id,omitempty"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Created  time.Time `json:"created"`
		Password string
		Role     Role `json:"role"`
		Stocks []Stocks `json:"stockList"`
	}

	Role struct {
		Id        uint32 `json:"_id,omitempty"`
		Authority string `json:"authority"`
	}

)

func encryptPassword(password []byte) string {
	h := sha256.New()
	h.Write(password)
	return fmt.Sprintf("%x", h.Sum(nil))

}

func GetUserByEmail(email string) (User, error) {
	user := User{}
	db := data.GetDatabase()
	defer db.Close()

	result := db.QueryRow("SELECT u.id, u.name, u.email, u.created, ur.id as roleId, ur.authority FROM user u left join user_role_connection uc on u.id = uc.user_id left join user_role ur on ur.id = uc.role_id WHERE u.email = ?", email)

	if err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Created, &user.Role.Id, &user.Role.Authority); err != nil {
		return user, err
	}

	return user, nil
}

func GetUserById(Id []uint8) (User, error) {
	user := User{}
	db := data.GetDatabase()
	defer db.Close()

	result := db.QueryRow("SELECT u.id, u.name, u.email, u.created, ur.id as roleId, ur.authority FROM user u left join user_role_connection uc on u.id = uc.user_id left join user_role ur on ur.id = uc.role_id WHERE u.id = ?", Id)

	if err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Created, &user.Role.Id, &user.Role.Authority); err != nil {
		return user, err
	}

	return user, nil
}

func (user *User) Create() (error) {
	db := data.GetDatabase()
	defer db.Close()

	//Inert User row
	stmt, err := db.Prepare("INSERT user SET name = ?, email = ?, password = ?, created = ?")

	if err != nil {
		return err
	}

	//Hash password

	hashedPassword := encryptPassword([]byte(user.Password))
	res, err := stmt.Exec(user.Name, user.Email, hashedPassword, time.Now())
	if err != nil {
		return err
	}

	insertedIndex, err := res.LastInsertId()

	if err != nil {
		return err
	}

	fmt.Printf("Inserted Success with index: %d\n", insertedIndex)

	//Inert user role connection
	stmt, err = db.Prepare("INSERT user_role_connection SET user_id = ?, role_id = ?")

	if err != nil {
		return err
	}

	res, err = stmt.Exec(insertedIndex, user.Role.Id)

	if err != nil {
		return err
	}

	insertedIndex, err = res.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Printf("Inserted User Role Connection Table. ID: %d\n", insertedIndex)

	return nil
}

func (user *User) Save() (error) {
	db := data.GetDatabase()
	defer db.Close()

	stmt, err := db.Prepare("update user set name = ? where id = ?")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(user.Name, user.Id)

	if err != nil {
		return err
	}

	if affect, err := res.RowsAffected(); err != nil {
		return err
	}else {
		fmt.Println(affect)
	}

	return nil
}

func GetRoleByAuthority(authority string) (role Role, err error) {
	db := data.GetDatabase()
	defer db.Close()

	row := db.QueryRow("SELECT id, authority FROM user_role WHERE authority = ?", authority)

	if err := row.Scan(&role.Id, &role.Authority); err != nil {
		return role, err
	}else {
		return role, nil
	}
}


/** Stock Function **/
func (this *User) SaveStock() error {
	db := data.GetDatabase()
	defer db.Close()

	stocks := this.Stocks

	for _, stock := range stocks {
		//Check if the stock already in stock list database
		var stockId int64 = -1
		err := db.QueryRow("SELECT id FROM stocks WHERE name = ? and symbol = ?", stock.Name, stock.Symbol).Scan(&stockId)

		if err != nil && err != sql.ErrNoRows {
			return err
		}

		fmt.Printf("Stock ID: %d\n", stockId)

		if stockId == -1 {
			stmt, err := db.Prepare("INSERT stocks SET name = ?, symbol = ?")

			if err != nil {
				return err
			}

			res, err := stmt.Exec(stock.Name, stock.Symbol)

			if err != nil {
				return err
			}

			stockId, err = res.LastInsertId()

			if err != nil {
				return err
			}

			log.Printf("Stock inserted to stock list. ID: %d\n", stockId)



		}

		stock.Id = uint32(stockId)
		//Connect between user and watch list -> user_stocks_connection
		var exists bool
		err = db.QueryRow("SELECT exists (SELECT id FROM user_stocks_connection WHERE user_id = ? and stock_id = ? and type = ?)", this.Id, stock.Id, stock.Type).Scan(&exists)

		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if !exists {
			stmt, err := db.Prepare("INSERT user_stocks_connection SET user_id = ?, stock_id = ?, type = ?")
			if err != nil {
				return err
			}
			fmt.Printf("Stock ID: %d", stock.Id)

			_, err = stmt.Exec(this.Id, stock.Id, stock.Type)

			if err != nil {
				log.Printf(err.Error())
				return err
			}
		}

	}
	return nil

}