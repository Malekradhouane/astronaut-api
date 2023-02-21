package errs

import "errors"

// ErrNoSuchEntity is thrown when no entity is found for the required key
var ErrNoSuchEntity = errors.New("no such entity")

// IsNoSuchEntityError return if the error is a wrapped (or not) not such entity error
func IsNoSuchEntityError(e error) bool {
	return errors.Is(e, ErrNoSuchEntity)
}
