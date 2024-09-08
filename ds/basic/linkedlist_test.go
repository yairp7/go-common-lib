package ds_test

import (
	"testing"

	ds "github.com/yairp7/go-common-lib/ds/basic"
	"gopkg.in/stretchr/testify.v1/assert"
)

func Test_Linkedlist_AddFront(t *testing.T) {
	list := ds.NewLinkedList[int]()
	list.AddFront(0)
	list.AddFront(1)
	list.AddFront(2)
	list.AddFront(3)
	assert.Equal(t, 4, list.Len())
	assert.Equal(t, list.PopHead().Data, 3)
	assert.Equal(t, list.PopTail().Data, 0)
}

func Test_Linkedlist_AddBack(t *testing.T) {
	list := ds.NewLinkedList[int]()
	list.AddBack(0)
	list.AddBack(1)
	list.AddBack(2)
	list.AddBack(3)
	assert.Equal(t, 4, list.Len())
	assert.Equal(t, list.PopHead().Data, 0)
	assert.Equal(t, list.PopTail().Data, 3)
}

func Test_Linkedlist_Remove(t *testing.T) {
	list := ds.NewLinkedList[int]()
	list.AddFront(0)
	e1 := list.AddFront(1)
	list.AddFront(2)
	e3 := list.AddFront(3)
	list.Remove(e1)
	list.Remove(e3)
	assert.Equal(t, 2, list.Len())
}

func Test_Linkedlist_At(t *testing.T) {
	list := ds.NewLinkedList[int]()
	e0 := list.AddFront(0)
	e1 := list.AddFront(1)
	e2 := list.AddFront(2)
	e3 := list.AddFront(3)
	list.AddFront(4)
	assert.Equal(t, e0, list.At(0))
	assert.Equal(t, e1, list.At(1))
	assert.Equal(t, e2, list.At(2))
	assert.Equal(t, e3, list.At(3))
}

func Test_Linkedlist_HeadTail(t *testing.T) {
	list := ds.NewLinkedList[int]()
	e0 := list.AddFront(0)
	list.AddFront(1)
	list.AddFront(2)
	e3 := list.AddFront(3)
	assert.Equal(t, e3, list.Head())
	assert.Equal(t, e0, list.Tail())
}

func Test_Linkedlist_PopTail(t *testing.T) {
	list := ds.NewLinkedList[int]()
	e0 := list.AddFront(0)
	e1 := list.AddFront(1)
	list.AddFront(2)
	list.AddFront(3)
	assert.Equal(t, e0, list.PopTail())
	assert.Equal(t, 3, list.Len())
	assert.Equal(t, e1, list.PopTail())
	assert.Equal(t, 2, list.Len())
}

func Test_Linkedlist_PopHead(t *testing.T) {
	list := ds.NewLinkedList[int]()
	list.AddFront(0)
	list.AddFront(1)
	e2 := list.AddFront(2)
	e3 := list.AddFront(3)
	assert.Equal(t, e3, list.PopHead())
	assert.Equal(t, 3, list.Len())
	assert.Equal(t, e2, list.PopHead())
	assert.Equal(t, 2, list.Len())
}

func Test_Linkedlist_Swap(t *testing.T) {
	list := ds.NewLinkedList[int]()
	list.AddFront(0)
	e1 := list.AddFront(1)
	list.AddFront(2)
	e3 := list.AddFront(3)
	list.Swap(e1, e3)
	assert.Equal(t, e1, list.Head())
	assert.Equal(t, e3, list.At(1))
	list.Swap(e1, e3)
	assert.Equal(t, e3, list.Head())
	assert.Equal(t, e1, list.At(1))
}
