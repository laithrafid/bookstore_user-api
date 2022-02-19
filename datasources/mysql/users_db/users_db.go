package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	logger "github.com/laithrafid/bookstore_user-api/utils/logger_utils"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var connErr error
	Client, connErr = sql.Open("mysql", dataSourceName)
	if connErr != nil {
		panic(connErr)
	}
	if connErr = Client.Ping(); connErr != nil {
		panic(connErr)
	}

	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")
}
