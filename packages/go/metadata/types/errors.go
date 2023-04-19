package types

import (
	"errors"
)

var (
	ErrDatastreamIdEmpty  = errors.New("datastream id cannot be empty")
	ErrDatastreamNotFound = errors.New("datastream not found")
)
