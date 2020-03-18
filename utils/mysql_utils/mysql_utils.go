package mysql_utils

import (
	"github.com/KestutisKazlauskas/go-users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {

		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("No record found.")
		}
		return errors.NewInternalServerError("error parsing database response.", err)
	}
	switch sqlErr.Number {
	case 1062: 
		return errors.NewBadRequestError("Invalid data.")
	}

	return errors.NewInternalServerError("error processing request.", err)
}