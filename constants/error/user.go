package error

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrPasswordIncorrect    = errors.New("password incorrect")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrUserAlreadyExists,
	ErrPasswordIncorrect,
	ErrPasswordDoesNotMatch,
}
