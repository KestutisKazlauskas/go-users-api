package mysql_utils

import (
	"github.com/KestutisKazlauskas/go-utils/rest_errors"
	"github.com/KestutisKazlauskas/go-utils/logger"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {

		if strings.Contains(err.Error(), errorNoRows) {
			return rest_errors.NewNotFoundError("No record found.")
		}
		return rest_errors.NewInternalServerError("error parsing database response.", err, logger.Log)
	}
	switch sqlErr.Number {
	case 1062: 
		return rest_errors.NewBadRequestError("Invalid data.")
	}

	return rest_errors.NewInternalServerError("error processing request.", err, logger.Log)
}