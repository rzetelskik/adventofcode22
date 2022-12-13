package util

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Heap[T constraints.Ordered] []T

func (h Heap[T]) Len() int {
	return len(h)
}

func (h Heap[T]) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *Heap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func MapSlice[T any, U any](xs []T, f func(T) U) []U {
	res := make([]U, 0)

	for _, x := range xs {
		res = append(res, f(x))
	}

	return res
}

func MapSliceWithError[T any, U any](xs []T, f func(T) (U, error)) ([]U, error) {
	res := make([]U, 0)

	for _, x := range xs {
		y, err := f(x)
		if err != nil {
			return []U{}, err
		}
		res = append(res, y)
	}

	return res, nil
}
