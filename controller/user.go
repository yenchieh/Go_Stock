package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go_stock_with_gin/common"
	"net/http"
	"errors"
	"github.com/go_stock_with_gin/model"
	"fmt"
	"database/sql"
)

func CreateUser(c *gin.Context){
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	if name == "" || email == "" || password == "" {
		common.RenderError(c, http.StatusBadRequest, errors.New("Name, Email and Password are required"), "Name, Email and Password are required")
		return
	}

	fmt.Printf("Name: %s, Email: %s, password: %s\n\n", name, email, password)

	role, err := model.GetRoleByAuthority("ROLE_USER")

	fmt.Printf("Role: %v\n", role.Id)
	if err != nil {
		common.RenderError(c, http.StatusInternalServerError, err, "Error on getting user Role")
		return
	}

	user := model.User{
		Name: name,
		Email: email,
		Password: password,
		Role: role,
	}

	if err := user.Create(); err != nil {
		panic(err);
	}

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