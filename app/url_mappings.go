package app

import (
	"github.com/KestutisKazlauskas/go-users-api/controllers/ping"
	"github.com/KestutisKazlauskas/go-users-api/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)

	//mapping user controllers
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/find", users.Find)
	router.POST("/users/login", users.Login)
}