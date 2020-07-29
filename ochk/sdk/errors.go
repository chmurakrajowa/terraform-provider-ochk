package sdk

import (
	"errors"
	"fmt"
)

type NotFoundError struct {
	Err error
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Resource not found: %s", e.Err)
}

func (e *NotFoundError) Unwrap() error { return e.Err }

func IsNotFoundError(err error) bool {
	var notFound *NotFoundError

	return errors.Is(err, notFound)
}
