package app

import (
	"github.com/gin-gonic/gin"
	"github.com/KestutisKazlauskas/go-utils/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {

	MapUrls()

	logger.Log.Info("About to start application")
	router.Run(":8081")	
}