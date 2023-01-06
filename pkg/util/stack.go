package util

type Stack[T any] []T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Pop() {
	if s.IsEmpty() {
		return
	}

	*s = (*s)[:len(*s)-1]
}

func (s *Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *Stack[T]) Top() T {
	return (*s)[len(*s)-1]
}
