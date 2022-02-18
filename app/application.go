package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/laithrafid/bookstore_user-api/utils/config_utils"
)

var (
	router = gin.Default()
)

func StartApplication() {
	config, err := config_utils.LoadConfig("../.")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	mapUrls()
	router.Run(config.ServerAddress)
}
