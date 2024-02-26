package error

import "errors"

var (
	ErrBadCredentials  = errors.New("email/password combination don't work")
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrUnknown         = errors.New("something went wrong")
	ErrEmail           = errors.New("email already register")
)
