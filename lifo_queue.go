package go_collections

import (
	"container/list"
	"sync"
)

type lifoQueue struct {
	data *list.List
	mut  *sync.RWMutex
}

// 后进先出队列
func NewLifoQueue() *lifoQueue {
	return &lifoQueue{data: list.New(), mut: new(sync.RWMutex)}
}

// 入队操作
func (q *lifoQueue) Put(v interface{}) {
	defer q.mut.Unlock()
	q.mut.Lock()
	q.data.PushFront(v)
}

// 出队操作
func (q *lifoQueue) Get() (interface{}, bool) {
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

// 返回队列长度
func (q *lifoQueue) Qsize() int {
	defer q.mut.RUnlock()
	q.mut.RLock()
	return q.data.Len()
}

// 判断队列是否为空
func (q *lifoQueue) IsEmpty() bool {
	defer q.mut.RUnlock()
	q.mut.RLock()
	return !(q.data.Len() > 0)
}
