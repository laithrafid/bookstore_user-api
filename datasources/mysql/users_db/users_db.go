package users_db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/laithrafid/bookstore_user-api/utils/config_utils"
)

var (
	Client *sql.DB
)

func init() {
	config, err := config_utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config of users_db:", err)
	}
	Client, connErr := sql.Open(config.DBDriver, config.DBSource)
	if connErr != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Successfuly connected to database")
}
