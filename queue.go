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

// 先进先出队列
type Queue struct {
	data *list.List
	mut  *sync.RWMutex
}

// 工厂函数 生成 `先进先出队列`
func NewQueue() *Queue {
	return &Queue{data: list.New(), mut: new(sync.RWMutex)}
}

// 入队操作
func (q *Queue) Put(v interface{}) {
	defer q.mut.Unlock()
	q.mut.Lock()
	q.data.PushFront(v)
}

// 出队操作
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

// 返回队列长度
func (q *Queue) Qsize() int {
	defer q.mut.RUnlock()
	q.mut.RLock()
	return q.data.Len()
}

// 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	defer q.mut.RUnlock()
	q.mut.RLock()
	return !(q.data.Len() > 0)
}
