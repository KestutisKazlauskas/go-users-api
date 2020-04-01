package services

import (
	"github.com/KestutisKazlauskas/go-users-api/domain/users"
	"github.com/KestutisKazlauskas/go-users-api/utils/date_utils"
	"github.com/KestutisKazlauskas/go-users-api/utils/crypto_utils"
	"github.com/KestutisKazlauskas/go-utils/rest_errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct {

}

type userServiceInterface interface {
	Get(int64) (*users.User, *rest_errors.RestErr)
	Create(users.User) (*users.User, *rest_errors.RestErr)
	Update(bool, users.User) (*users.User, *rest_errors.RestErr) 
	Delete(int64) *rest_errors.RestErr
	Find(string) (users.Users, *rest_errors.RestErr)
	Login(users.LoginRequest) (*users.User, *rest_errors.RestErr) 
}

func (service *userService) Get(userId int64) (*users.User, *rest_errors.RestErr) {
	
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (service *userService) Create(user users.User) (*users.User, *rest_errors.RestErr) {

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

func (service *userService) Update(isParial bool, user users.User) (*users.User, *rest_errors.RestErr) {
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

func (service *userService) Delete(userId int64) *rest_errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (service *userService) Find(status string) (users.Users, *rest_errors.RestErr) {
	dao := &users.User{}

	return dao.FindByStatus(status)
}

func (service *userService) Login(request users.LoginRequest) (*users.User, *rest_errors.RestErr) {
	user := &users.User{
		Email: request.Email,
	}

	if err := user.FindByEmail(); err != nil {
		return nil, err
	}

	if err := user.ValidatePassowrd(request.Password); err != nil {
		return nil, err
	}

	return user, nil
}