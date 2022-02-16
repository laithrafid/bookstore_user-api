package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	Users "github.com/laithrafid/bookstore_user-api/domain/users"
	"github.com/laithrafid/bookstore_user-api/services"
	"github.com/laithrafid/bookstore_user-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user Users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveerr := services.CreateUser(user)
	if saveerr != nil {
		c.JSON(saveerr.Status, saveerr)
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
