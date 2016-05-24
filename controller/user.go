package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go_stock_with_gin/common"
	"net/http"
	"github.com/go_stock_with_gin/model"
	"fmt"
	"database/sql"
	"github.com/go_stock_with_gin/data"
	"errors"
)

type UserCredential struct {
	Email string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Name string `form:"name" json:"name"`
}

func RegisterOrLogin(c *gin.Context) {
	var json UserCredential

	err := c.BindJSON(&json)
	if err != nil {
		common.RenderError(c, http.StatusBadRequest, err, "Email and Password are required")
		return
	}

	fmt.Printf("Name: %s, Email: %s, password: %s\n", json.Name, json.Email, json.Password)

	user, err := model.GetUserByEmail(json.Email)

	if err != nil && err != sql.ErrNoRows {
		common.RenderError(c, http.StatusInternalServerError, err, "Error on getting User")
		return
	}

	if err == sql.ErrNoRows {
		user, err = createUser(json.Email, json.Password, json.Name)
		if err != nil {
			common.RenderError(c, http.StatusInternalServerError, err, "Error on save user")
			return;
		}
	} else {
		//Doing login
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": user,
	})
}

func createUser(email, password, name string) (user model.User, err error){
	role, err := model.GetRoleByAuthority("ROLE_USER")

	fmt.Printf("Role: %v\n", role.Id)
	if err != nil {
		return user, errors.New("Error on getting User Role")
	}

	user = model.User{
		Name: name,
		Email: email,
		Password: password,
		Role: role,
	}

	return user, user.Create()
}

func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := model.GetUserByEmail(email)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "The user not found",
		})
		return;
	}else{
		common.RenderError(c, http.StatusInternalServerError, err, "Error on retreving user")
		return;
	}

	fmt.Println(user.Role.Authority)

	fmt.Println(user)
}

func CheckEmail(c *gin.Context){
	email := c.Param("email")

	db := data.GetDatabase()
	defer db.Close()

	user := db.QueryRow("SELECT * FROM user WHERE email = ?", email)

	err := user.Scan()

	if err != nil && err == sql.ErrNoRows{
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"exist": false,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"exist": true,
		})
	}

}
