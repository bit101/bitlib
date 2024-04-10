// Package collections contains custom collection types.
package collections

// MakeGrid creates a fully allocated 2d slice of a given type and size.
func MakeGrid[T any](rows, cols int) [][]T {
	arr := make([][]T, rows)
	for i := 0; i < rows; i++ {
		arr[i] = make([]T, cols)
	}
	return arr
}
