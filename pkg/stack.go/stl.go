package stl

import (
	"errors"
)

type Stack[T any] struct {
	stack []T
}

type Queue[T any] struct {
	queue []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack[T]) Push(value T) {
	s.stack = append(s.stack, value)

}

func (s *Stack[T]) Top() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.stack[len(s.stack)-1], nil
}

func (s *Stack[T]) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.stack)
}
