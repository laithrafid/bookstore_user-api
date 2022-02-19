package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "root", "127.0.0.1:3306", "users_db")
	// config, err := config_utils.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("cannot load config of users_db:", err)
	// }
	var err error
	Client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		println("connection problem")
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Successfuly connected to database")
}
