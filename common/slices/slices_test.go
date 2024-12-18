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

func Test_Filter(t *testing.T) {
	t.Run("FilterNils", func(t *testing.T) {
		srcVals := []float64{0.523, 24.53, 66.1, 637.3, 84.2}
		src := make([]*float64, 10)
		for i := range src {
			if i%2 == 0 {
				src[i] = &srcVals[i/2]
			} else {
				src[i] = nil
			}
		}

		srcWithoutNils := slices.FilterNils(src)

		for i := range srcWithoutNils {
			assert.Equal(t, srcVals[i], *srcWithoutNils[i])
		}
	})
}
