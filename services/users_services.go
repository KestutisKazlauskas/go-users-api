package services

import (
	"github.com/KestutisKazlauskas/go-users-api/domain/users"
	"github.com/KestutisKazlauskas/go-users-api/utils/date_utils"
	"github.com/KestutisKazlauskas/go-users-api/utils/crypto_utils"
	"github.com/KestutisKazlauskas/go-users-api/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct {

}

type userServiceInterface interface {
	Get(int64) (*users.User, *errors.RestErr)
	Create(users.User) (*users.User, *errors.RestErr)
	Update(bool, users.User) (*users.User, *errors.RestErr) 
	Delete(int64) *errors.RestErr
	Find(string) (users.Users, *errors.RestErr)
}

func (service *userService) Get(userId int64) (*users.User, *errors.RestErr) {
	
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (service *userService) Create(user users.User) (*users.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	//bussnes logic goes here
	user.CreatedAt = date_utils.GetNowDBFormat()
	user.Status = users.StatusActive

	password, err := crypto_utils.GetBCryptHash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = password

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil

}

func (service *userService) Update(isParial bool, user users.User) (*users.User, *errors.RestErr) {
	current := &users.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}

	if isParial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Validate(); err != nil {
		return nil, err
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (service *userService) Delete(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (service *userService) Find(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}

	return dao.FindByStatus(status)
}