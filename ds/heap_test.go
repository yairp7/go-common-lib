package ds_test

import (
	"slices"
	"sort"
	"testing"

	ds "github.com/yairp7/go-common-lib/ds"
	"gopkg.in/stretchr/testify.v1/assert"
)

func Test_MinHeap(t *testing.T) {
	values := []int{53, 22, 33, 65, 12, 1, 3, 77}
	minHeap := ds.NewMinHeap[int]()
	for _, v := range values {
		minHeap.Push(&ds.HeapNode[int]{
			Weight: float64(v),
			Data:   v,
		})
	}

	sort.Ints(values)
	i := 0
	for minHeap.Len() > 0 {
		v := minHeap.Pop()
		assert.Equal(t, values[i], v.Data)
		i++
	}
}

func Test_MaxHeap(t *testing.T) {
	values := []int{53, 22, 33, 65, 12, 1, 3, 77}
	maxHeap := ds.NewMaxHeap[int]()
	for _, v := range values {
		maxHeap.Push(&ds.HeapNode[int]{
			Weight: float64(v),
			Data:   v,
		})
	}

	sort.Ints(values)
	slices.Reverse(values)
	i := 0
	for maxHeap.Len() > 0 {
		v := maxHeap.Pop()
		assert.Equal(t, values[i], v.Data)
		i++
	}
}
