package users

// date transfer object , from here you can write logic to transfer from db to my app
import (
	"strings"

	"github.com/laithrafid/bookstore_utils-go/errors_utils"
)

const (
	StatusActive  = "active"
	StatusPending = "pending"
	StatusBanned  = "banned"
	StatusCreated = "created"
	StatusRemoved = "removed"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user *User) Validate() errors_utils.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors_utils.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors_utils.NewBadRequestError("invalid password")
	}
	//TODO: password policy validator function
	return nil
}
