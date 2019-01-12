package collections

import (
	"container/heap"
	"sync"
)

type PriorityQueue struct {
	nodes []*PqNode
	mut   *sync.RWMutex
}

// PriorityQueue Node
type PqNode struct {
	Value           string
	Priority, index int
}

func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{mut: new(sync.RWMutex)}
	heap.Init(pq)
	return pq
}

func (pq *PriorityQueue) Put(v *PqNode) {
	defer pq.mut.Unlock()
	pq.mut.Lock()
	heap.Push(pq, v)
}

func (pq *PriorityQueue) Get() (interface{}, bool) {
	defer pq.mut.Unlock()
	pq.mut.Lock()
	if len(pq.nodes) > 0 {
		item := heap.Pop(pq)
		return item, true
	}
	return nil, false
}

func (pq PriorityQueue) Qsize() int {
	defer pq.mut.RUnlock()
	pq.mut.RLock()
	return len(pq.nodes)
}

func (pq *PriorityQueue) IsEmpty() bool {
	defer pq.mut.RUnlock()
	pq.mut.RLock()
	return !(len(pq.nodes) > 0)
}

// `Sort` interface Len()
func (pq PriorityQueue) Len() int {
	return len(pq.nodes)
}

// `Sort` interface Less()
func (pq PriorityQueue) Less(i, j int) bool {
	return pq.nodes[i].Priority > pq.nodes[j].Priority
}

// `Sort` interface Swap()
func (pq PriorityQueue) Swap(i, j int) {
	pq.nodes[i], pq.nodes[j] = pq.nodes[j], pq.nodes[i]
	pq.nodes[i].index, pq.nodes[j].index = i, j
}

// `Heap` interface Push()
func (pq *PriorityQueue) Push(v interface{}) {
	item := v.(*PqNode)
	item.index = len(pq.nodes)
	pq.nodes = append(pq.nodes, item)
	heap.Fix(pq, item.index)
}

// `Heap` interface Pop()
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old.nodes)
	item := old.nodes[n-1]
	item.index = -1
	pq.nodes = old.nodes[0 : n-1]
	return item
}
