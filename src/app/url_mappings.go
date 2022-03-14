package app

import (
	"github.com/laithrafid/user-api/src/controllers/ping"
	"github.com/laithrafid/user-api/src/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
	// internal access only
	router.POST("/users", users.Create)
	router.POST("/users/login", users.Login)
}
