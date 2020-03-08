package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/KestutisKazlauskas/go-users-api/domain/users"
	"github.com/KestutisKazlauskas/go-users-api/services"
	"github.com/KestutisKazlauskas/go-users-api/utils/errors"
)

func GetUser(controller *gin.Context) {
	userId, userErr := strconv.ParseInt(controller.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.NewBadRequestError("invalid user_id")
		controller.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		controller.JSON(getErr.Status, getErr)
		return 
	}

	controller.JSON(http.StatusOK, user)
}

func CreateUser(controller *gin.Context) {
	var user users.User
	if err := controller.ShouldBindJSON(&user); err != nil {
		// ShouldbindJSON do the json validation with controller.Request.Body
		/*	Do this code for us:
			bytes, err := ioutil.ReadAll(controller.Request.Body)
			if err != nil {
				return
			}
			if err := json.Unmarshal(bytes, &user); err != nil {
				fmt.Println(err.Error())

				return
			}
		*/
		restErr := errors.NewBadRequestError("invalid JSON body")
		controller.JSON(restErr.Status, restErr)
		return 
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		controller.JSON(saveErr.Status, saveErr)
		return 
	}
	controller.JSON(http.StatusCreated, result)
}

func FindUser(controller *gin.Context) {
	controller.String(http.StatusNotImplemented, "Need some work!")
}