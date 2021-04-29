package helpers

import (
	"errors"
	"time"
)

var (
	// ErrChannel variable holds default error message if no channel is provided.
	ErrChannel = errors.New("target channel or id can not be empty")
	// TimeValue holds current date and time in unix format.
	TimeValue = "⏰ " + time.Now().Format(time.UnixDate)
)
