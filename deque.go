package collections

import (
	"container/list"
	"sync"
)

type Deque struct {
	data *list.List
	mut  *sync.RWMutex
}

func NewDeque() *Deque {
	return &Deque{data: list.New(), mut: new(sync.RWMutex)}
}

func (q *Deque) PutLeft(v interface{}) {
	defer q.mut.Unlock()
	q.mut.Lock()
	q.data.PushFront(v)
}

func (q *Deque) PutRight(v interface{}) {
	defer q.mut.Unlock()
	q.mut.Lock()
	q.data.PushBack(v)
}

func (q *Deque) GetLeft() (interface{}, bool) {
	defer q.mut.Unlock()
	q.mut.Lock()
	if q.data.Len() > 0 {
		iter := q.data.Front()
		v := iter.Value
		q.data.Remove(iter)
		return v, true
	}
	return nil, false
}

func (q *Deque) GetRight() (interface{}, bool) {
	defer q.mut.Unlock()
	q.mut.Lock()
	if q.data.Len() > 0 {
		iter := q.data.Back()
		v := iter.Value
		q.data.Remove(iter)
		return v, true
	}
	return nil, false
}

func (q *Deque) Qsize() int {
	defer q.mut.RUnlock()
	q.mut.RLock()
	return q.data.Len()
}

func (q *Deque) IsEmpty() bool {
	defer q.mut.RUnlock()
	q.mut.RLock()
	return !(q.data.Len() > 0)
}
