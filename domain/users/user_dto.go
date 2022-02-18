package Dusers

// date transfer object , from here you can write logic to transfer from db to my app
import (
	"strings"

	"github.com/laithrafid/bookstore_user-api/utils/errors_utils"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors_utils.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors_utils.NewBadRequestError("invalid email address")
	}
	return nil
}
