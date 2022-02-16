package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	Users "github.com/laithrafid/bookstore_user-api/domain/users"
	"github.com/laithrafid/bookstore_user-api/services"
)

func CreateUser(c *gin.Context) {
	var user Users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		//TODO : Return Bad request to the caller
		return
	}
	result, saveerr := services.CreateUser(user)
	if saveerr != nil {
		//TODO: Handle User Creation Error
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "implement me")

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
