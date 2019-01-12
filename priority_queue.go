package collections

import (
	"container/heap"
	"sync"
)

// 优先队列
type PriorityQueue struct {
	nodes []*PqNode
	mut   *sync.RWMutex
}

// 优先队列节点
type PqNode struct {
	Value    string
	Priority int
	index    int
}

// 工厂函数 生成 `优先队列`
func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{mut: new(sync.RWMutex)}
	heap.Init(pq)
	return pq
}

// 入队操作
func (pq *PriorityQueue) Put(v *PqNode) {
	defer pq.mut.Unlock()
	pq.mut.Lock()
	heap.Push(pq, v)
}

// 出队操作
func (pq *PriorityQueue) Get() (interface{}, bool) {
	defer pq.mut.Unlock()
	pq.mut.Lock()
	if len(pq.nodes) > 0 {
		item := heap.Pop(pq)
		return item, true
	}
	return nil, false
}

// 返回队列长度
func (pq PriorityQueue) Qsize() int {
	defer pq.mut.RUnlock()
	pq.mut.RLock()
	return len(pq.nodes)
}

// 判断队列是否为空
func (pq *PriorityQueue) IsEmpty() bool {
	defer pq.mut.RUnlock()
	pq.mut.RLock()
	return !(len(pq.nodes) > 0)
}

// `Sort` 接口 Len()
func (pq PriorityQueue) Len() int {
	return len(pq.nodes)
}

// `Sort` 接口 Less()
func (pq PriorityQueue) Less(i, j int) bool {
	return pq.nodes[i].Priority > pq.nodes[j].Priority
}

// `Sort` 接口 Swap()
func (pq PriorityQueue) Swap(i, j int) {
	pq.nodes[i], pq.nodes[j] = pq.nodes[j], pq.nodes[i]
	pq.nodes[i].index, pq.nodes[j].index = i, j
}

// `Heap` 接口 Push()
func (pq *PriorityQueue) Push(v interface{}) {
	item := v.(*PqNode)
	item.index = len(pq.nodes)
	pq.nodes = append(pq.nodes, item)
	heap.Fix(pq, item.index)
}

// `Heap` 接口 Pop()
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old.nodes)
	item := old.nodes[n-1]
	item.index = -1
	pq.nodes = old.nodes[0 : n-1]
	return item
}
