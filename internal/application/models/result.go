package models

type Result[T any] struct {
	Value           T
	Error           error
	IsInternalError bool
	HttpStatusCode  int
}

func NewResultSuccess[T any](value T) *Result[T] {
	return &Result[T]{Value: value}
}

func NewResultFailure[T any](error error, isInternalError bool, httpStatusCode int) *Result[T] {
	return &Result[T]{Error: error, IsInternalError: isInternalError, HttpStatusCode: httpStatusCode}
}
