package services

import (
	Dusers "github.com/laithrafid/bookstore_user-api/domain/users"
	"github.com/laithrafid/bookstore_user-api/utils/errors_utils"
)

func CreateUser(user Dusers.User) (*Dusers.User, *errors_utils.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
func GetUser(userId int64) (*Dusers.User, *errors_utils.RestErr) {
	result := &Dusers.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}
