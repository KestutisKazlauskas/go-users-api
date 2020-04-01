package users

import (
	"strings"
	"github.com/KestutisKazlauskas/go-utils/rest_errors"
	"github.com/KestutisKazlauskas/go-users-api/utils/crypto_utils"
)

const (
	StatusActive="active"
)

type User struct {
	Id 			int64  `json:"id"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email 		string `json:"email"`
	CreatedAt 	string `json:"created_at"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user *User) Validate() *rest_errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return rest_errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)

	if user.Password == "" {
		return rest_errors.NewBadRequestError("invalid password address")
	}

	return nil
}

func (user *User) ValidatePassowrd(password string) *rest_errors.RestErr {
	isPasswordCorrect := crypto_utils.CheckBCryptPassword(user.Password, password)

	if !isPasswordCorrect {
		return rest_errors.NewBadRequestError("Incorrect logins") 
	}

	return nil 
}