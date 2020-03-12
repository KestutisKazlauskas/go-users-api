package users
// access to database for user code of database only herer

import (
	"github.com/KestutisKazlauskas/go-users-api/utils/errors"
	"github.com/KestutisKazlauskas/go-users-api/utils/date_utils"
	"github.com/KestutisKazlauskas/go-users-api/datasources/mysql/users_db"
	"fmt"
	"strings"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, created_at) VALUES (?, ?, ?, ?)"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Clinet.Ping(); err != nil {
		panic(err)
	}
	result := userDB[user.Id]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	//What is this??
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedAt = result.CreatedAt

	return nil
}

func (user *User) Save() *errors.RestErr {

	statment, err := users_db.Clinet.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	//important to close connection after code execution
	defer statment.Close()
	user.CreatedAt = date_utils.GetNowString()
	insertResult, err := statment.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt)

	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("Error saving user %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("Error on LastInsertedId %s", err.Error()))
	}

	user.Id = userId
	return nil
}

