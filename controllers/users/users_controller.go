package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(controller *gin.Context) {
	controller.String(http.StatusNotImplemented, "Need some work!")
}

func CreateUser(controller *gin.Context) {
	controller.String(http.StatusNotImplemented, "Need some work!")
}

func FindUser(controller *gin.Context) {
	controller.String(http.StatusNotImplemented, "Need some work!")
}