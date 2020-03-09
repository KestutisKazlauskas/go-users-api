package users
// access to database for user code of database only herer

import (
	"github.com/KestutisKazlauskas/go-users-api/utils/errors"
	"github.com/KestutisKazlauskas/go-users-api/utils/date"
	"fmt"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
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
	current := userDB[user.Id]
	if current != nil {
		return errors.NewBadRequestError("User already exists")
	}
	
	user.CreatedAt = date.GetNowString()

	userDB[user.Id] = user
	return nil
}

