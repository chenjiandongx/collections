package collections

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var maxCnt int = 10e4

func yieldRandomArray(cnt int) []int {
	res := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		res[i] = rand.Int()
	}
	return res
}

func assertSort(items []int) bool {
	for i := 0; i < len(items)-1; i++ {
		if items[i] > items[i+1] {
			return false
		}
	}
	return true
}

type StdItems struct {
	data []int
}

func (o StdItems) Less(i, j int) bool {
	return o.data[i] < o.data[j]
}

func (o StdItems) Swap(i, j int) {
	o.data[i], o.data[j] = o.data[j], o.data[i]
}

func (o StdItems) Len() int {
	return len(o.data)
}

func TestBubbleSort(t *testing.T) {
	items := yieldRandomArray(maxCnt)
	BubbleSort(items)
	assert.True(t, assertSort(items))
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		items := yieldRandomArray(maxCnt)
		BubbleSort(items)
	}
}

func TestInsertSort(t *testing.T) {
	items := yieldRandomArray(maxCnt)
	InsertSort(items)
	assert.True(t, assertSort(items))
}

func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		items := yieldRandomArray(maxCnt)
		InsertSort(items)
	}
}

func TestStdSort(t *testing.T) {
	items := StdItems{yieldRandomArray(maxCnt)}
	sort.Sort(items)
	assert.True(t, assertSort(items.data))
}

func BenchmarkStdSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Sort(StdItems{yieldRandomArray(maxCnt)})
	}
}

func TestQuickSort(t *testing.T) {
	items := yieldRandomArray(maxCnt)
	QuickSort(items)
	assert.True(t, assertSort(items))
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		items := yieldRandomArray(maxCnt)
		QuickSort(items)
	}
}

func TestShellSort(t *testing.T) {
	items := yieldRandomArray(maxCnt)
	ShellSort(items)
	assert.True(t, assertSort(items))
}

func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		items := yieldRandomArray(maxCnt)
		ShellSort(items)
	}
}

func TestMergeSort(t *testing.T) {
	items := yieldRandomArray(maxCnt)
	MergeSort(items)
	assert.True(t, assertSort(items))
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		items := yieldRandomArray(maxCnt)
		MergeSort(items)
	}
}
