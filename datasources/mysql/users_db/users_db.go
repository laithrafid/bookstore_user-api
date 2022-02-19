package users_db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/laithrafid/bookstore_user-api/utils/config_utils"
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
)

func init() {
	config, err := config_utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config of variables for db:", err)
	}
	username := config.Username
	password := config.Password
	host := config.Host
	schema := config.Schema

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var connErr error
	Client, connErr = sql.Open("mysql", dataSourceName)
	if connErr != nil {
		panic(connErr)
	}
	if connErr = Client.Ping(); err != nil {
		panic(connErr)
	}

	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")
}
