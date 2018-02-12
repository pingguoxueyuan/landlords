package logic

import (
	"errors"
)

var (
	ErrSeatIsUsed        = errors.New("seat is used")
	ErrInvalidSeatStatus = errors.New("invalid seat status")
	ErrInvalidParameter  = errors.New("invalid parameter")
	ErrReachMaxPlayer    = errors.New("reach max player num")
	ErrServerBusy        = errors.New("server is busy")
	ErrNotFoundPlayer    = errors.New("not found player")
)
