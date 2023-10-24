package fga

import (
	"errors"
)

var (
	// ErrFGAMissingHost is returned when a host is not provided
	ErrFGAMissingHost = errors.New("invalid OpenFGA config: missing host")
)
