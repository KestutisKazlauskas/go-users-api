package crypto_utils

import (
	"github.com/KestutisKazlauskas/go-utils/rest_errors"
	"github.com/KestutisKazlauskas/go-utils/logger"
	"golang.org/x/crypto/bcrypt"
)

func GetBCryptHash(input string) (string, *rest_errors.RestErr) {
	byteInput := []byte(input)

	hash, err := bcrypt.GenerateFromPassword(byteInput, bcrypt.MinCost)
	if err != nil {
		return "", rest_errors.NewInternalServerError("Error hashing password", err, logger.Log)
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