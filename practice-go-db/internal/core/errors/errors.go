package errors

import "errors"

var (
	ErrorFailToAddDepartment = errors.New("fail to add department")
	ErrorNoSuchDepartment = errors.New("no such department")
	ErrorNothingToGet = errors.New("no result")

	ErrorInvalidId = errors.New("invalid id value")
	ErrorInvalidLinit = errors.New("invalid linit value")
	ErrorInvalidOffset = errors.New("invalid offset value")
)
