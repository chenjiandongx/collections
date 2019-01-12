package collections

import (
	"container/list"
	"sync"
)

/*
	RWMutex 的使用主要事项
	1. 读锁的时候无需等待读锁的结束
	2. 读锁的时候要等待写锁的结束
	3. 写锁的时候要等待读锁的结束
	4. 写锁的时候要等待写锁的结束

	RWMutex 的四种操作方法
	RLock()		//读锁定
	RUnlock()	//读解锁
	Lock()		//写锁定
	Unlock()	//写解锁
*/

type Queue struct {
	data *list.List
	mut  *sync.RWMutex
}

func NewQueue() *Queue {
	return &Queue{data: list.New(), mut: new(sync.RWMutex)}
}

func (q *Queue) Put(v interface{}) {
	defer q.mut.Unlock()
	q.mut.Lock()
	q.data.PushFront(v)
}

func (q *Queue) Get() (interface{}, bool) {
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

func (q *Queue) Qsize() int {
	defer q.mut.RUnlock()
	q.mut.RLock()
	return q.data.Len()
}

func (q *Queue) IsEmpty() bool {
	defer q.mut.RUnlock()
	q.mut.RLock()
	return !(q.data.Len() > 0)
}
