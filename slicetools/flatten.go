package slicetools

func Flatten[S ~[][]T, T any](values S) []T {
	var result []T
	for _, seq := range values {
		result = append(result, seq...)
	}
	return result
}
