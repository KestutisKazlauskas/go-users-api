package app

import (
	"github.com/gin-gonic/gin"
	"github.com/KestutisKazlauskas/go-users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {

	MapUrls()

	logger.Info("About to start application")
	router.Run(":8081")
	
}