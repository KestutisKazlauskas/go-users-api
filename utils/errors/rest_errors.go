package errors
import (
	"net/http"
)

type RestErr struct {
	Message string  `json:"message"`
	Status 	int 	`json:"status"`
	Error 	string 	`json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	error := RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad_request",
	}

	return &error
}

func NewNotFoundError(message string) *RestErr {
	error := RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "bad_request",
	}

	return &error
}