package slices

import (
	"errors"
	"slices"
)

type Mat[T any] struct {
	Values        []T
	Rows, Columns int
}

func NewMat[T any](rows, columns int) *Mat[T] {
	return &Mat[T]{
		Rows:    rows,
		Columns: columns,
		Values:  make([]T, rows*columns),
	}
}

func (m *Mat[T]) Get(row, col int, defaultValue T) T {
	if (row >= m.Rows || row < 0) || (col >= m.Columns || col < 0) {
		return defaultValue
	}
	return m.Values[row*m.Columns+col]
}

func (m *Mat[T]) Set(row, col int, val T) error {
	if (row >= m.Rows || row < 0) || (col >= m.Columns || col < 0) {
		return errors.New("index out of bounds")
	}
	m.Values[row*m.Columns+col] = val
	return nil
}

func CreateMat[T any](rows, columns int, create func(row, col int) T) [][]T {
	matrix := make([][]T, rows)
	for row := 0; row < rows; row++ {
		matrix[row] = make([]T, columns)
		for col := 0; col < columns; col++ {
			matrix[row][col] = create(row, col)
		}
	}
	return matrix
}

func ForEachMat[T any](s [][]T, f func(row, col int, val T)) {
	if len(s) == 0 {
		return
	}

	for row := 0; row < len(s); row++ {
		for col := 0; col < len(s[0]); col++ {
			f(row, col, s[row][col])
		}
	}
}

func CopyMat[T any](s [][]T) [][]T {
	n := len(s)
	m := len(s[0])
	clone := make([][]T, n)
	data := make([]T, n*m)

	for row := 0; row < n; row++ {
		copy(data[row*m:row*m+m], s[row])
		clone[row] = data[row*m : row*m+m]
	}

	return clone
}

func Flat[T any](s [][]T) []T {
	if s == nil || s[0] == nil {
		return nil
	}

	rows := len(s)
	columns := len(s[0])
	size := rows * columns

	flatSlice := make([]T, size)
	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			flatSlice[row*columns+col] = s[row][col]
		}
	}

	return flatSlice
}

func FlatMap[T any, R any](s [][]T, mapFunc func(item T) R) []R {
	if s == nil || s[0] == nil {
		return nil
	}

	if len(s) == 0 {
		return nil
	}

	rows := len(s)
	columns := len(s[0])
	size := rows * columns
	flatSlice := make([]R, size)

	for row := 0; row < rows; row++ {
		flatSlice = slices.Replace(flatSlice, row*columns, row*columns+columns, Map(s[row], mapFunc)...)
	}

	return flatSlice
}
