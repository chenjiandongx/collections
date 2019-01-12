package collections

import (
	"container/list"
	"sync"
)

type LifoQueue struct {
	data *list.List
	mut  *sync.RWMutex
}

func NewLifoQueue() *LifoQueue {
	return &LifoQueue{data: list.New(), mut: new(sync.RWMutex)}
}

func (q *LifoQueue) Put(v interface{}) {
	defer q.mut.Unlock()
	q.mut.Lock()
	q.data.PushFront(v)
}

func (q *LifoQueue) Get() (interface{}, bool) {
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

func (q *LifoQueue) Qsize() int {
	defer q.mut.RUnlock()
	q.mut.RLock()
	return q.data.Len()
}

func (q *LifoQueue) IsEmpty() bool {
	defer q.mut.RUnlock()
	q.mut.RLock()
	return !(q.data.Len() > 0)
}
