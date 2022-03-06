package users_db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/laithrafid/bookstore_utils-go/config_utils"
	"github.com/laithrafid/bookstore_utils-go/logger_utils"
)

var (
	Client *sql.DB
)

func init() {
	config, err := config_utils.LoadConfig(".")
	if err != nil {
		logger_utils.Error("cannot load mysql config:", err)
	}
	driver := config.MysqlDBDriver
	source := config.MysqlDBSource

	var connErr error
	Client, connErr = sql.Open(driver, source)
	if connErr != nil {
		logger_utils.Error("Error trying to connect to database", connErr)
	}
	if connErr = Client.Ping(); connErr != nil {
		logger_utils.Error("Error trying to ping to database", connErr)
	}

	mysql.SetLogger(logger_utils.GetLogger())
	logger_utils.Info("successfully connected to backend Database")
}
