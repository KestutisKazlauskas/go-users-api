package app

import (
	"github.com/KestutisKazlauskas/go-users-api/controllers/ping"
	"github.com/KestutisKazlauskas/go-users-api/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)

	//mapping user controllers
	//router.GET("/users/find", users.FindUser)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}