package ds

import "sync"

type Queue[T any] struct {
	size         int
	items        []T
	isThreadSafe bool
	lock         sync.Mutex
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		size:  0,
		items: []T{},
	}
}

func NewQueueAsync[T any]() *Queue[T] {
	return &Queue[T]{
		size:         0,
		items:        []T{},
		isThreadSafe: true,
	}
}

func (q *Queue[T]) Push(item T) {
	if q.isThreadSafe {
		q.lock.Lock()
		defer q.lock.Unlock()
	}

	if q.size < len(q.items) {
		q.items[q.size] = item
	} else {
		q.items = append(q.items, item)
	}
	q.size++
}

func (q *Queue[T]) Pop() T {
	if q.isThreadSafe {
		q.lock.Lock()
		defer q.lock.Unlock()
	}

	item := q.items[0]
	q.items = q.items[1:]
	q.size--
	return item
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Clear() {
	clear(q.items)
	q.size = 0
}
