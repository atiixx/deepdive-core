package stack

import (
	"github.com/atiixx/deepdive/core/list"
)

type Stack[T comparable] struct {
	Values list.List[T]
}

func (s *Stack[T]) Push(value T) {
	s.Values.Prepend(value)
}

func (s *Stack[T]) Pop() T {
	if s.Values.Size() == 0 {
		var zero T
		return zero
	}
	return s.Values.RemoveAt(0)
}

func (s Stack[T]) Peek() T {
	if s.Values.Size() == 0 {
		var zero T
		return zero
	}
	return s.Values.Head.Value
}

func (s Stack[T]) String() string {
	return s.Values.String()
}
