package users_db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/laithrafid/bookstore_user-api/utils/config_utils"
	logger "github.com/laithrafid/bookstore_user-api/utils/logger_utils"
)

var (
	Client *sql.DB
)

func init() {
	config, err := config_utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load mysql config:", err)
	}
	driver := config.DBDriver
	source := config.DBSource

	var connErr error
	Client, connErr = sql.Open(driver, source)
	if connErr != nil {
		panic(connErr)
	}
	if connErr = Client.Ping(); connErr != nil {
		panic(connErr)
	}

	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")
}
