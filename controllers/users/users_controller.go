package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/KestutisKazlauskas/go-users-api/domain/users"
	"github.com/KestutisKazlauskas/go-users-api/services"
	"github.com/KestutisKazlauskas/go-utils/rest_errors"
	"github.com/KestutisKazlauskas/go-oauth/oauth"
)

func getUserId(userIdParam string) (int64, *rest_errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("invalid user_id")
	}

	return userId, nil
}

func Get(context *gin.Context) {
	//TODO fix authentication for Public Ip if not authenicatd it should be public
	//TODO return error if bad authentication token is send.
	if err := oauth.Authenticate(context.Request); err != nil {
		context.JSON(err.Status, err)
		return
	}
	userId, idErr := getUserId(context.Param("user_id"))
	if idErr != nil {
		context.JSON(idErr.Status, idErr)
		return
	}
	user, getErr := services.UserService.Get(userId)
	if getErr != nil {
		context.JSON(getErr.Status, getErr)
		return 
	}

	if oauth.GetUserId(context.Request) == user.Id {
		context.JSON(http.StatusOK, user.Marshall(false))
		return 
	}

	context.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(context.Request)))
}

func Create(context *gin.Context) {
	var user users.User
	if err := context.ShouldBindJSON(&user); err != nil {
		// ShouldbindJSON do the json validation with context.Request.Body
		/*	Do this code for us:
			bytes, err := ioutil.ReadAll(context.Request.Body)
			if err != nil {
				return
			}
			if err := json.Unmarshal(bytes, &user); err != nil {
				fmt.Println(err.Error())

				return
			}
		*/
		restErr := rest_errors.NewBadRequestError("invalid JSON body")
		context.JSON(restErr.Status, restErr)
		return 
	}
	result, saveErr := services.UserService.Create(user)
	if saveErr != nil {
		context.JSON(saveErr.Status, saveErr)
		return 
	}
	context.JSON(http.StatusCreated, result.Marshall(context.GetHeader("X-Public") == "true"))
}

func Update(context *gin.Context) {
	userId, idErr := getUserId(context.Param("user_id"))
	if idErr != nil {
		context.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := context.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid JSON body")
		context.JSON(restErr.Status, restErr)
		return 
	}

	user.Id = userId
	isPartial := context.Request.Method == http.MethodPatch

	result, err :=services.UserService.Update(isPartial, user)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	context.JSON(http.StatusOK, result.Marshall(context.GetHeader("X-Public") == "true"))
}

func Delete(context *gin.Context) {
	userId, idErr := getUserId(context.Param("user_id"))
	if idErr != nil {
		context.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UserService.Delete(userId); err != nil {
		context.JSON(err.Status, err)
		return
	}

	context.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

func Find(context *gin.Context) {
	status := context.Query("status")
	
	users, err := services.UserService.Find(status)
	if err != nil {
		context.JSON(err.Status, err)
		return 
	}

	context.JSON(http.StatusOK, users.Marshall(context.GetHeader("X-Public") == "true"))
}

func Login(context *gin.Context) {
	var request users.LoginRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json")
		context.JSON(restErr.Status, restErr)
		return
	}

	user, err := services.UserService.Login(request)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	context.JSON(http.StatusOK, user.Marshall(false))
}

