package services

import (
	Users "github.com/laithrafid/bookstore_user-api/domain/users"
	"github.com/laithrafid/bookstore_user-api/utils/errors"
)

func CreateUser(user Users.User) (*Users.User, *errors.RestErr) {
	return &user, nil
}
