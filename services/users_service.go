package services

import Users "github.com/laithrafid/bookstore_user-api/domain/users"

func CreateUser(user Users.User) (*Users.User, error) {
	return &user, nil
}
