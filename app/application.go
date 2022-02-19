package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/laithrafid/bookstore_user-api/utils/config_utils"
	logger "github.com/laithrafid/bookstore_user-api/utils/logger_utils"
)

var (
	router = gin.Default()
)

func StartApplication() {
	config, err := config_utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config of application:", err)
	}
	mapUrls()
	logger.Info("starting the application ....")
	router.Run(config.ServerAddress)
}
