package slices

func Index[T comparable](slice []T, t T) int {
	for i, v := range slice {
		if t == v {
			return i
		}
	}
	return -1
}

func Shift[T any](slice []T) (T, []T) { return slice[0], slice[1:] }

func Unshift[T any](slice []T, t T) []T { return append([]T{t}, slice...) }
