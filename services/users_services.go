package services

import (
	"github.com/KestutisKazlauskas/go-users-api/domain/users"
	"github.com/KestutisKazlauskas/go-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	return &user, nil

}