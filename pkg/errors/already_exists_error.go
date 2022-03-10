package errors

import (
	"errors"
	"fmt"
)

type alreadyExistsError struct {
	errorMessage error
}

func (e alreadyExistsError) Error() string {
	return fmt.Sprintf("%v", e.errorMessage)
}

func IsAlreadyExistsError(err error) bool {
	_, ok := err.(alreadyExistsError)
	return ok
}

func AlreadyExists(errorMessage string) alreadyExistsError {
	return alreadyExistsError{errorMessage: errors.New(errorMessage)}
}
