package repositories

import "errors"

var (
	errInvalidModel  = errors.New("invalid model")
	errAlreadyExists = errors.New("user already exists")

	errPageOrLimitAreWrong = errors.New("page or limit cannot be zero")
)