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
