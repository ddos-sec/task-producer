package task_producer

import "errors"

var (
	ErrCountInvalid        = errors.New("count must be greater than 0")
	ErrDurationInvalid     = errors.New("duration must be greater than 0")
	ErrMakeTaskFuncInvalid = errors.New("makeTaskFunc must not be nil")
)
