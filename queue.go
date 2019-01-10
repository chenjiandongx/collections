package go_collections

import (
	"container/heap"
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
type queue struct {
	data *list.List
	mut  *sync.RWMutex
}

// 先进先出队列
func NewQueue() *queue {
	return &queue{data: list.New(), mut: new(sync.RWMutex)}
}

// 入队操作
func (q *queue) Put(v interface{}) {
	defer q.mut.Unlock()
	q.mut.Lock()
	q.put(v)
}

// 出队操作
func (q *queue) Get() (bool, interface{}) {
	defer q.mut.Unlock()
	q.mut.Lock()
	if q.qsize() > 0 {
		iter := q.get()
		v := iter.Value
		q.data.Remove(iter)
		return true, v
	}
	return false, nil
}

// 队列长度
func (q *queue) Qsize() int {
	defer q.mut.RUnlock()
	q.mut.RLock()
	return q.qsize()
}

// 队列是否为空
func (q *queue) IsEmpty() bool {
	defer q.mut.RUnlock()
	q.mut.RLock()
	return !(q.qsize() > 0)
}

func (q *queue) qsize() int {
	return q.data.Len()
}

func (q *queue) get() *list.Element {
	return q.data.Back()
}

func (q *queue) put(v interface{}) {
	q.data.PushFront(v)
}

type lifoQueue queue

// 后进先出队列
func NewLifoQueue() *lifoQueue {
	return &lifoQueue{data: list.New(), mut: new(sync.RWMutex)}
}

func (lq *lifoQueue) get() *list.Element {
	return lq.data.Front()
}

func (lq *lifoQueue) put(v interface{}) {
	lq.data.PushFront(v)
}

type PqNode struct {
	Value    string
	Priority int
	index    int
}

type none struct{}

type priorityQueue struct {
	nodes []*PqNode
	mut   *sync.RWMutex
}

// 优先级队列
func NewPriorityQueue() *priorityQueue {
	pq := &priorityQueue{mut: new(sync.RWMutex)}
	heap.Init(pq)
	return pq
}

// 入队操作
func (pq *priorityQueue) Put(v interface{}) {
	heap.Push(pq, v)
}

// 出队操作
func (pq *priorityQueue) Get() (bool, interface{}) {
	item := heap.Pop(pq)
	switch item.(type) {
	case none:
		return false, nil
	default:
		return true, item
	}
}

// 队列长度
func (pq priorityQueue) Qsize() int {
	defer pq.mut.RUnlock()
	pq.mut.RLock()
	return pq.qsize()
}

// 队列是否为空
func (pq *priorityQueue) IsEmpty() bool {
	defer pq.mut.RUnlock()
	pq.mut.RLock()
	return !(pq.qsize() > 0)
}

// `Sort` 接口 Len()
func (pq priorityQueue) Len() int {
	return len(pq.nodes)
}

// `Sort` 接口 Less()
func (pq priorityQueue) Less(i, j int) bool {
	return pq.nodes[i].Priority > pq.nodes[j].Priority
}

// `Sort` 接口 Swap()
func (pq priorityQueue) Swap(i, j int) {
	pq.nodes[i], pq.nodes[j] = pq.nodes[j], pq.nodes[i]
	pq.nodes[i].index, pq.nodes[j].index = i, j
}

// `Heap` 接口 Push()
func (pq *priorityQueue) Push(v interface{}) {
	defer pq.mut.Unlock()
	pq.mut.Lock()
	item := v.(*PqNode)
	item.index = pq.qsize()
	pq.nodes = append(pq.nodes, item)
	heap.Fix(pq, item.index)
}

// `Heap` 接口 Pop()
func (pq *priorityQueue) Pop() interface{} {
	defer pq.mut.Unlock()
	pq.mut.Lock()
	if pq.qsize() > 0 {
		old := *pq
		n := len(old.nodes)
		item := old.nodes[n-1]
		item.index = -1
		pq.nodes = old.nodes[0 : n-1]
		return item
	}
	return none{}
}

func (pq *priorityQueue) qsize() int {
	return len(pq.nodes)
}
