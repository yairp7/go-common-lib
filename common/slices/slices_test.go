package slices_test

import (
	"strings"
	"testing"

	"github.com/yairp7/go-common-lib/common/slices"
	"gopkg.in/stretchr/testify.v1/assert"
)

var theStrings = []string{"sgsdg", "sdfgdsf", "sdgdsg", "sgrgwef", "Sdgqef"}
var theMap = map[string]string{
	"sgsdg":   "sgsdg",
	"sdfgdsf": "sdfgdsf",
	"sdgdsg":  "sdgdsg",
	"sgrgwef": "sgrgwef",
	"Sdgqef":  "Sdgqef",
}

func TestSome(t *testing.T) {
	strs := []string{"asfa", "234", "af4", "sff"}

	result := slices.Some(strs, func(item string) bool {
		return strings.Contains(item, "34")
	})
	assert.True(t, result)

	result = slices.Some(strs, func(item string) bool {
		return strings.Contains(item, "3f34")
	})
	assert.False(t, result)
}

func BenchmarkForEach(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slices.ForEach(theStrings, func(item string) {})
	}
}

func BenchmarkFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slices.Filter(theStrings, func(item string) bool {
			return item == "sdfgdsf"
		})
	}
}

func BenchmarkReduce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slices.Reduce(theStrings, 0, func(acc int, item string) int {
			return acc + len(item)
		})
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slices.Map(theStrings, func(item string) string {
			return item
		})
	}
}

func BenchmarkKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slices.Keys(theMap)
	}
}

func BenchmarkValues(b *testing.B) {
	m := map[string]string{
		"sgsdg":   "sgsdg",
		"sdfgdsf": "sdfgdsf",
		"sdgdsg":  "sdgdsg",
		"sgrgwef": "sgrgwef",
		"Sdgqef":  "Sdgqef",
	}

	for i := 0; i < b.N; i++ {
		slices.Values(m)
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

func BenchmarkCopy2D(b *testing.B) {
	matrix := generateMatrix()
	for i := 0; i < b.N; i++ {
		slices.Copy2D(matrix)
	}
}
