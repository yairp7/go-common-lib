package slices_test

import (
	"testing"

	"github.com/yairp7/go-common-lib/common/slices"
)

var theStrings = []string{"sgsdg", "sdfgdsf", "sdgdsg", "sgrgwef", "Sdgqef"}
var theMap = map[string]string{
	"sgsdg":   "sgsdg",
	"sdfgdsf": "sdfgdsf",
	"sdgdsg":  "sdgdsg",
	"sgrgwef": "sgrgwef",
	"Sdgqef":  "Sdgqef",
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
