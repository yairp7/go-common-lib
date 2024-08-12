package slices_test

import (
	"testing"

	"github.com/yairp7/go-common-lib/common/slices"
	"gopkg.in/stretchr/testify.v1/assert"
)

func TestFlat(t *testing.T) {
	mat := [][]int{
		{5, 5, 5},
		{6, 6, 6},
		{7, 7, 7},
	}

	expectedMat := []int{5, 5, 5, 6, 6, 6, 7, 7, 7}

	actualMat := slices.Flat(mat)

	assert.Equal(t, expectedMat, actualMat)
}

func TestFlatMap(t *testing.T) {
	mat := [][]int{
		{5, 5, 5},
		{6, 6, 6},
		{7, 7, 7},
	}

	expectedMat := []int{6, 6, 6, 7, 7, 7, 8, 8, 8}

	actualMat := slices.FlatMap(mat, func(item int) int {
		return item + 1
	})

	assert.Equal(t, expectedMat, actualMat)
}

func BenchmarkFlatMap(b *testing.B) {
	mat := [][]int{
		{5, 5, 5},
		{6, 6, 6},
		{7, 7, 7},
	}

	for i := 0; i < b.N; i++ {
		slices.FlatMap(mat, func(item int) int {
			return item + 1
		})
	}
}

func generateMatrix() [][]int {
	matrix := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		matrix[i] = make([]int, 1000)
		for j := 0; j < 1000; j++ {
			matrix[i][j] = i * j
		}
	}
	return matrix
}

func TestCopyMat(t *testing.T) {
	mat := generateMatrix()
	copiedMat := slices.CopyMat(mat)
	assert.Equal(t, mat, copiedMat)
}

func BenchmarkCopyMat(b *testing.B) {
	matrix := generateMatrix()
	for i := 0; i < b.N; i++ {
		slices.CopyMat(matrix)
	}
}

func TestMat(t *testing.T) {
	mat := slices.NewMat[int](50, 50)
	mat.Set(5, 8, 77)
	assert.Equal(t, mat.Get(5, 8, -1), 77)
}
