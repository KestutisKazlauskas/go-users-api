package crypto_utils

import (
	"github.com/KestutisKazlauskas/go-users-api/utils/errors"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func GetBCryptHash(input string) (string, *errors.RestErr) {
	byteInput := []byte(input)

	hash, err := bcrypt.GenerateFromPassword(byteInput, bcrypt.MinCost)
	if err != nil {
		return "", errors.NewInternalServerError(fmt.Sprintf("Error hashing password %s", err.Error()))
	}

	return string(hash), nil
}