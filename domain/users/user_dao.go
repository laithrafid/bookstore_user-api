package Users

// date access object  , here you can write from/to and query dbs
import (
	"fmt"

	"github.com/laithrafid/bookstore_user-api/utils/errors_utils"

	"github.com/laithrafid/bookstore_user-api/datasources/mysql/users_db"
	"github.com/laithrafid/bookstore_user-api/utils/date_utils"
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
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors_utils.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors_utils.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	user.DateCreated = date_utils.GetNowString()

	usersDB[user.Id] = user
	return nil
}
