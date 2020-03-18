package crypto_utils

import (
	"github.com/KestutisKazlauskas/go-users-api/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

func GetBCryptHash(input string) (string, *errors.RestErr) {
	byteInput := []byte(input)

	hash, err := bcrypt.GenerateFromPassword(byteInput, bcrypt.MinCost)
	if err != nil {
		return "", errors.NewInternalServerError("Error hashing password", err)
	}

	return string(hash), nil
}

func CheckBCryptPassword(hash string, input string) bool {
	byteHash := []byte(hash)
	byteInput := []byte(input)

	err := bcrypt.CompareHashAndPassword(byteHash, byteInput)

	if err != nil {
		return false
	}

	return true
}