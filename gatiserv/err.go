package gatiserv

import "errors"

var (
	// ErrUnkonow - unknow error
	ErrUnkonow = errors.New("unknow error")

	// ErrInvalidPlayerState - invalid PlayerState
	ErrInvalidPlayerState = errors.New("invalid PlayerState")
)
