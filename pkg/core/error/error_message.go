package error

import "errors"

var (
	ErrBadCredentials  = errors.New("Email/password combination don't work")
	ErrUnauthenticated = errors.New("Unauthenticated")
	ErrUnknown         = errors.New("Something went wrong")
	ErrEmail           = errors.New("Email already register")
)
