package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(controller *gin.Context) {
	controller.String(http.StatusOK, "pong")
}