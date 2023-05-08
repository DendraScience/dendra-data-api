package types

import (
	"errors"
)

var (
	ErrConvertFromEmpty = errors.New("convert from cannot be inferred or is empty")
	ErrConvertToEmpty   = errors.New("convert to cannot be empty")
)
