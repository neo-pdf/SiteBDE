package bcrypt

import (
	"errors"
)

var ErrHashTooShort = errors.New("crypto/bcrypt: hashedSecret too short to be a bcrypted password")
var ErrMismatchedHashAndPassword = errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password")
var ErrPasswordTooLong = errors.New("bcrypt: password length exceeds 72 bytes")

const (
	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)
