package model

import (
	"github.com/go_stock_with_gin/data"
	"fmt"
	"time"
	"crypto/sha256"
)

type (
	User struct {
		Id       []uint8 `json:"_id,omitempty"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Created  time.Time `json:"created"`
		Password string
		Role     Role `json:"role"`
	}

	Role struct {
		Id        []uint8 `json:"_id,omitempty"`
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

	fmt.Printf("Inserted Success with index: %d", insertedIndex)

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
	fmt.Printf("Inserted User Role Connection Table. ID: %d", insertedIndex)

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