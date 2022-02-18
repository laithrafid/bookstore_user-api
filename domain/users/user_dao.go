package Dusers

// date access object  , here you can write from/to and query dbs
import (
	"fmt"

	"github.com/laithrafid/bookstore_user-api/utils/errors_utils"

	"github.com/laithrafid/bookstore_user-api/datasources/mysql/users_db"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT FROM users(first_name, last_name, email, date_created) WHERE Id=useris;"
	queryGetUsers   = "SELECT * FROM users_db.users;"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors_utils.RestErr {
	result := usersDB[user.Id]
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	if result == nil {
		return errors_utils.NewNotFoundError(fmt.Sprintf("user %d not found in  db", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}
func (user *User) Save() *errors_utils.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	instertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors_utils.NewInternalServerError(
			fmt.Sprintf("error when trying to Save user: %s", err.Error()))
	}
	userId, err := instertResult.LastInsertId()
	if err != nil {
		return errors_utils.NewInternalServerError(fmt.Sprintf("error when trying to Save user with Id: %s", err.Error()))
	}
	user.Id = userId
	return nil
}
