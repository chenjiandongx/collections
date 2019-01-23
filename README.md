# 📂 collctions

> Golang 实现的 collections 模块，灵感来自 [Python queue](https://docs.python.org/3/library/queue.html) 和 [Python collections](https://docs.python.org/3/library/collections.html)

[![Build Status](https://travis-ci.org/chenjiandongx/collections.svg?branch=master)](https://travis-ci.org/chenjiandongx/collections) [![Go Report Card](https://goreportcard.com/badge/github.com/chenjiandongx/collections)](https://goreportcard.com/report/github.com/chenjiandongx/collections) [![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT) [![GoDoc](https://godoc.org/github.com/chenjiandongx/collections?status.svg)](https://godoc.org/github.com/chenjiandongx/collections)

## 📚 目录
* [Queue - 先进先出队列](#Queue)
* [LifoQueue - 后进先出队列](#LifoQueue)
* [PriorityQueue - 优先队列](#PriorityQueue)
* [Deque - 双端队列](#Deque)
* [OrderedMap - 有序 Map](#OrderedMap)
* [Counter - 计数器](#Counter)
* [AVLTree - AVL 树](#AVLTree)
* [Sort - 排序](#Sort)

### 🔰 安装&引用
```bash
$ go get github.com/chenjiandongx/collections

import "github.com/chenjiandongx/collections"
```

### 📦 Collections
### Queue
> 先进先出队列（线程安全）

📝 方法集
```shell
Get()(interface{}, bool)    // 出队
Put(v interface{})          // 入队
Qsize() int                 // 返回队列长度
IsEmpty() bool              // 判断队列是否为空
```

✏️ 示例
```go
var nums = 1000

q := collections.NewQueue()
var item interface{}
var ok bool
for i := 0; i < nums; i++ {
    q.Put(i)
}
for i := 0; i < nums; i++ {
    if item, ok = q.Get(); ok {
        fmt.Println(item.(int))
    }
}

fmt.Println(q.IsEmpty())
fmt.Println(q.Qsize())
```

### LifoQueue
> 后进先出队列（线程安全）

📝 方法集
```shell
Get()(interface{}, bool)    // 出队
Put(v interface{})          // 入队
Qsize() int                 // 返回队列长度
IsEmpty() bool              // 判断队列是否为空
```

✏️ 示例
```go
var nums = 1000

q := collections.NewLifoQueue()
var item interface{}
var ok bool
for i := 0; i < nums; i++ {
    q.Put(i)
}
for i := nums-1; i >=0; i-- {
    if item, ok = q.Get(); ok {
        fmt.Println(item.(int))
    }
}

fmt.Println(q.IsEmpty())
fmt.Println(q.Qsize())
```

### PriorityQueue
> 优先队列（线程安全）

📝 方法集
```shell
Get()(interface{}, bool)    // 出队
Put(v *PqNode)              // 入队
Qsize() int                 // 返回队列长度
IsEmpty() bool              // 判断队列是否为空

// 优先队列节点
type PqNode struct {
    Value           string
    Priority, index int
}
```

✏️ 示例
```go
var nums = 1000

q := collections.NewPriorityQueue()

for i := 0; i < nums; i++ {
    r := rand.Int()
    q.Put(&collections.PqNode{Value: string(r), Priority: rand.Int()})
}

for i := 0; i < nums/2; i++ {
    item1, _ := q.Get()
    item2, _ := q.Get()
    fmt.Println(item1.(*collections.PqNode).Priority > item2.(*collections.PqNode).Priority)
}

fmt.Println(q.IsEmpty())
fmt.Println(q.Qsize())
```

### Deque
> 双端队列（线程安全）

📝 方法集
```shell
GetLeft()(interface{}, bool)        // 左边出队
GetRight()(interface{}, bool)       // 右边出队
PutLeft(v interface{})              // 左边入队
PutRight(v interface{})             // 右边入队
Qsize() int                         // 返回队列长度
IsEmpty() bool                      // 判断队列是否为空
```

✏️ 示例
```go
var nums = 1000
q := collections.NewDeque()

var item interface{}
var ok bool

for i := 0; i < nums; i++ {
    q.PutLeft(i)
}
fmt.Println(q.Qsize())

for i := nums - 1; i >= 0; i-- {
    q.PutRight(i)
}
fmt.Println(q.Qsize())

for i := 0; i < nums; i++ {
    item, ok = q.GetRight()
    fmt.Println(item, ok)
}
for i := nums - 1; i >= 0; i-- {
    item, ok = q.GetLeft()
    fmt.Println(item, ok)
}

item, ok = q.GetLeft()
fmt.Println(item, ok)

item, ok = q.GetRight()
fmt.Println(item, ok)
```

### OrderedMap
> 有序 Map，接口设计参考 [cevaris/ordered_map](https://github.com/cevaris/ordered_map)

📝 方法集
```shell
Set(key, value interface{})                 // 新增键值对
Get(key interface{}) (interface{}, bool)    // 取值
Delete(key interface{}) bool                // 删除键
Iter() (interface{}, interface{}, bool)     // 遍历
Len() int                                   // 键值对数量
// 指针回退到 Head，遍历时 current 指针会向后移动 BackToHead 使其移动到头指针，以便下一次从头遍历
BackToHead()                               
```

✏️ 示例
```go
maxNum := 100
om := collections.NewOrderedMap()
for i := 0; i < maxNum; i++ {
    om.Set(i, i+1)
}

fmt.Println(om.Len())
om.Delete(0)
fmt.Println(om.Len())

for k, v, ok := om.Iter(); ok; k, v, ok = om.Iter() {
    fmt.Println(k, v)
}

om.BackToHead()
for k, v, ok := om.Iter(); ok; k, v, ok = om.Iter() {
    fmt.Println(k, v)
}
```

📣 讨论

有序 Map 在 Golang 中应该是一个十分常见的需求，Map 最大的优势就是它的查找性能，**理论上** Map 查找的时间复杂度是常数级的。但实际情况如何，我们可以通过 benchmark 来验证。在 [Go Maps Don’t Appear to be O(1)](https://medium.com/@ConnorPeet/go-maps-are-not-o-1-91c1e61110bf) 这篇文章中，作者测试了 Golang Map 查找的实际性能，不过作者是基于 Go1.4 的，版本有点旧了。下面是我修改了作者的测试案例后在 Go1.10 下跑出来的结果。

![](https://user-images.githubusercontent.com/19553554/51075377-83f8cd80-16c5-11e9-9973-4904a4661aeb.png)

上图是使用 [go-echarts](https://github.com/chenjiandongx/go-echarts) 绘制的。测试是通过与二分查找来对比的，二分查找的时间复杂度为 **O(log2n)**。很明显，在 10e5 数量级下两者的性能差别还不是特别大，主要差距是在 10e6 后体现的。结论：Map 的性能优于 **O(log2n)**，但不是常数级。

**collections.OrderdMap 🆚 cevaris/ordered_map**

本来我一直使用的是 [cevaris/ordered_map](https://github.com/cevaris/ordered_map)，后来自己重新实现了一个。实现完就与它进行了性能测试对比。它是基于两个 Map 实现的，而我是使用的 Map+LinkedList，LinkedList 在删除和插入操作上的时间复杂度都是 **O(1)**，用其来存储 Map key 的顺序是一个很好的选择。

同样的测试代码，BenchMark 结果如下
```shell
goos: windows
goarch: amd64
pkg: github.com/chenjiandongx/collections
BenchmarkCollectionsSet-8        2000000               689 ns/op             187 B/op          3 allocs/op
BenchmarkCevarisSet-8            1000000              1212 ns/op             334 B/op          3 allocs/op
BenchmarkCollectionsGet-8        2000000               823 ns/op             187 B/op          3 allocs/op
BenchmarkCevarisGet-8            1000000              1281 ns/op             334 B/op          3 allocs/op
BenchmarkCollectionsIter-8       2000000               670 ns/op             187 B/op          3 allocs/op
BenchmarkCevarisIter-8           1000000              1341 ns/op             366 B/op          4 allocs/op
```
**collections.OrderedMap Win 🖖 性能+内存占用全部占优 🚀**

### Counter
> 计数器

📝 方法集
```shell
// key-value item
type Item struct {
    k interface{}
    v int
}

Add(keys ...interface{})            // 新增 item
Get(key interface{}) int            // 获取 key 计数
GetAll() []Item                     // 获取全部 key 计数
Top(n int) []Item                   // 获取前 key 计数
Delete(key interface{}) bool        // 删除 key，成功返回 true，key 不存在返回 false
Len() int                           // key 数量
```

✏️ 示例
```go
c := collections.NewCounter()
c.Add("a", "b", "c", "d", "a", "c")
fmt.Println(c.Get("A"))
fmt.Println(c.Get("a"))
fmt.Println(c.Get("b"))
fmt.Println(c.Top(2))
fmt.Println(c.Len())
fmt.Println(c.All())
c.Delete("a")
```

### AVLTree
> AVL 二叉自平衡查找树

📝 方法集
```shell
NewAVLTree() *AVLTree       // 生成 AVL 树
Insert(v int)               // 插入节点
Search(v int) bool          // 搜索节点
Delete(v int) bool          // 删除节点
GetMaxValue() int           // 获取所有节点中的最大值
GetMinValue() int           // 获取所有节点中的最小值
AllValues() []int           // 返回排序后所有值
```

✏️ 示例
```go
var maxNum = 100

tree := NewAVLTree()
for i := 0; i < maxNum; i++ {
    tree.Insert(i)
    tree.Insert(maxNum + i)
}
fmt.Println(len(tree.AllValues()))
fmt.Println(tree.GetMaxValue())
fmt.Println(tree.GetMinValue())
fmt.Println(tree.Search(50))
fmt.Println(tree.Search(100))
fmt.Println(tree.Search(-10))
fmt.Println(tree.Delete(-10))
fmt.Println(tree.Delete(10))
```

📣 讨论

AVL 树是自平衡树的一种，其通过左旋和右旋来调整自身的平衡性，使其左右子树的高度差最大不超过 1。AVL 在插入、查找、删除的平时时间复杂度都是 O(logn)，在基本的 BST（二叉查找树）中，理想情况的效率也是为 O(logn)，但由于操作的性能其实是依赖于树的高度，而 BST 最坏的情况会导致树退化成链表，此时时间复杂度就变为 O(n)，为了解决这个问题，自平衡二叉树应运而生。

AVL 的主要精髓在于`旋转`，旋转分为 4 种情况，左旋，左旋+右旋，右旋，右旋+左旋。调整树结构后需要重新计算树高。

**左子树左节点失衡**
> 左左情况 直接右旋
```shell
    x                
  x        => 右旋         x
x                       x    x
```

**左子树右节点失衡**
> 左右情况 先左旋后右旋
```shell
  x                        x     
x         => 左旋         x       => 右旋        x
  x                     x                     x    x
```

**右子树右节点失衡**
> 右右情况 直接左旋
```shell
x                
  x       => 左旋          x
    x                   x    x
```

**右子树左节点失衡**
> 右左情况 先右旋后左旋
```shell
x                      x     
  x       => 右旋        x       => 左旋        x
x                          x                 x    x
```

AVL 主要的性能消耗主要在插入，因为其需要通过旋转来维护树的平衡，但如果使用场景是经常需要排序和查找数据的话，AVL 还是可以展现其良好的性能的。

**benchmark**
```
BenchmarkAVLInsert10e1-6        2000000000               0.00 ns/op
BenchmarkAVLInsert10e2-6        2000000000               0.00 ns/op
BenchmarkAVLInsert10e3-6        2000000000               0.00 ns/op
BenchmarkAVLInsert10e4-6        2000000000               0.02 ns/op
BenchmarkAVLInsert10e5-6        1000000000               0.82 ns/op
BenchmarkAVLSearch-6            2000000000               0.00 ns/op
BenchmarkAVLDelete-6            2000000000               0.00 ns/op
```

### Sort

📝 方法集
```shell
BubbleSort()        // 冒泡排序
InsertionSort()     // 插入排序
QuickSort()         // 快速排序
ShellSort()         // 希尔排序
HeapSort()          // 堆排序
MergeSort()         // 归并排序
```

✏️ 示例
```go
var maxCnt = 10e4

func yieldRandomArray() []int {
    res := make([]int, maxCnt)
    for i := 0; i < maxCnt; i++ {
        res[i] = rand.Int()
    }
    return res
}

BubbleSort(yieldRandomArray())
InsertionSort(yieldRandomArray())
QuickSort(yieldRandomArray())
ShellSort(yieldRandomArray())
HeapSort(yieldRandomArray())
MergeSort(yieldRandomArray())
```

📣 讨论

**排序算法时间复杂度比较**

| 排序算法 |  是否稳定  |  平均    |   最好  |    最差   |   动画演示  |
| -------- | --------- |----------| --------| -------- | ----------- |
| BubbleSort | 是 | O(n^2) |  O(n) |  O(n^2) | ![](https://upload.wikimedia.org/wikipedia/commons/3/37/Bubble_sort_animation.gif) |
| InsertionSort | 是 | O(n^2) |  O(n) |  O(n^2) | ![](https://upload.wikimedia.org/wikipedia/commons/2/25/Insertion_sort_animation.gif) |
| QuickSort | 否 | O(nlogn) | O(nlogn) |  O(n^2) | ![](https://upload.wikimedia.org/wikipedia/commons/6/6a/Sorting_quicksort_anim.gif) |
| ShellSort | 否 |O(nlogn) |  O(n) | O(n^2)  | ![](https://upload.wikimedia.org/wikipedia/commons/2/25/Insertion_sort_animation.gif) |
| HeapSort | 否 | O(nlogn) |  O(nlogn) | O(nlogn) | ![](https://upload.wikimedia.org/wikipedia/commons/1/1b/Sorting_heapsort_anim.gif) |
| MergeSort | 是 | O(nlogn) |  O(nlogn) | O(nlogn) | ![](https://upload.wikimedia.org/wikipedia/commons/c/c5/Merge_sort_animation2.gif) |

通过 benchmark 来测试平均排序性能

**数据随机分布**
```go
var maxCnt int = 10e4

func yieldRandomArray(cnt int) []int {
    res := make([]int, cnt)
    for i := 0; i < cnt; i++ {
        res[i] = rand.Int()
    }
    return res
}
```

运行结果
```shell
BenchmarkBubbleSort-8                  1        17361549400 ns/op
BenchmarkInsertionSort-8               1        1934826900 ns/op
BenchmarkQuickSort-8                 100          10651807 ns/op
BenchmarkShellSort-8                 100          16476199 ns/op
BenchmarkHeapSort-8                  100          14231607 ns/op
BenchmarkMergeSort-8                 100          14840583 ns/op
```

冒泡和直接插入排序在随机数据集的排序性能最差，为 O(n^2)，剩余 4 种排序快排效率最佳，其他 3 者性能很接近。

**换两种极端的数据分布方式**

**数据升序分布**
```go
func yieldArrayAsce(cnt int) []int {
    res := make([]int, cnt)
    for i := 0; i < cnt; i++ {
        res[i] = i
    }
    return res
}
```

运行结果
```shell
BenchmarkBubbleSort-8               5000            266690 ns/op
BenchmarkInsertionSort-8           10000            213429 ns/op
BenchmarkQuickSort-8                   1        3291222900 ns/op
BenchmarkShellSort-8                1000           1716406 ns/op
BenchmarkHeapSort-8                  200           6806788 ns/op
BenchmarkMergeSort-8                 300           4677485 ns/op
```

在数据基本升序的情况下，冒泡和直接插入排序能够取得良好的性能。而快排就给跪了，就是最差的 O(n^2) 了。

**数据降序分布**
```go
func yieldArrayDesc(cnt int) []int {
    res := make([]int, cnt)
    for i := 0; i < cnt; i++ {
        res[i] = cnt-i
    }
    return res
}
```

运行结果
```shell
BenchmarkBubbleSort-8                  1        6710048800 ns/op
BenchmarkInsertionSort-8               1        3881599100 ns/op
BenchmarkQuickSort-8                   1        3373971200 ns/op
BenchmarkShellSort-8                 500           2876371 ns/op
BenchmarkHeapSort-8                  200           7081150 ns/op
BenchmarkMergeSort-8                 300           4448222 ns/op
```

在数据基本降序的情况下，冒泡和直接插入排序一如既往的差，快排又给跪了，又是 O(n^2)...

那自己写的排序和 Golang 官方提供的 sort.Sort 排序方法对比，效率如何呢


定义一个 struct，实现 sort.Interface
```go
import "sort"

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
```

只取 n(logn) 复杂度的排序算法与标准 sort 进行对比

**数据随机分布**
```shell
BenchmarkStdSort-8                            50          22978524 ns/op
BenchmarkQuickSort-8                         100          11648689 ns/op
BenchmarkShellSort-8                         100          17353544 ns/op
BenchmarkHeapSort-8                          100          14501199 ns/op
BenchmarkMergeSort-8                         100          13793086 ns/op
```

是不是眼前一亮 😂，自己写的快排居然这么厉害，比标准的 sort 快了不止两倍？？？ 这里出现这样的情况的主要原因是 sort 实现了 sort.Interface，该接口需要有三个方法 Less()/Len()/Swap()，而接口的类型转换是有成本的。**通用**意味着**牺牲**，这是**专**和**精**权衡后的结果。当然，标准的 sort 大部分情况的性能都是可以接受的，也是比较方便的。但当你需要追求极致性能的话，自己针对特定需求实现排序算法肯定会是更好的选择。

**数据升序分布**
```shell
BenchmarkStdSort-8                           200           7285511 ns/op
BenchmarkQuickSort-8                           1        3351046900 ns/op
BenchmarkShellSort-8                        1000           1679506 ns/op
BenchmarkHeapSort-8                          200           6632256 ns/op
BenchmarkMergeSort-8                         300           4308582 ns/op
```

是不是又是眼前一亮 🤣，我去 为什么这次标准的排序比快排快了这么多，官方的排序不也是快排吗？（这个测试结果看起来好像也没人会比快排慢是吧 😅）

**数据降序分布**
```shell
BenchmarkStdSort-8                           200           7405331 ns/op
BenchmarkQuickSort-8                           1        3390954400 ns/op
BenchmarkShellSort-8                         500           2900240 ns/op
BenchmarkHeapSort-8                          200           7091124 ns/op
BenchmarkMergeSort-8                         300           4295169 ns/op
```

emmmmmmm，同上 😓

关于官方排序的具体实现，可以参考 [src/sort/sort.go](https://golang.org/src/sort/sort.go)，实际上是直接插入排序，快速排序，堆排序和归并排序的组合排序。[这篇文章](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example/blob/master/chapter03/03.1.md) 对这部分有介绍

最后，按官方的排序针对自己想要的数据类型排序 不使用接口那套 会是什么效率呢 对比上面排序中最快的算法以及接口实现的 sort

**数据随机分布**
```shell
BenchmarkStdSort-8                           100          22649399 ns/op
BenchmarkQuickSort-8                         100          10870924 ns/op
BenchmarkStdSortWithoutInterface-8           100          10511605 ns/op
```

**数据升序分布**
```shell
BenchmarkStdSort-8                           200           7006117 ns/op
BenchmarkShellSort-8                        1000           1667537 ns/op
BenchmarkStdSortWithoutInterface-8          1000           1619643 ns/op
```

**数据降序分布**
```shell
BenchmarkStdSort-8                           200           7614625 ns/op
BenchmarkShellSort-8                         500           3051834 ns/op
BenchmarkStdSortWithoutInterface-8          1000           1689479 ns/op
```

🖖 [Sort](https://github.com/chenjiandongx/collections/blob/master/std_sort.go) 完胜！！！

故事到这里还没有结束，我们还可以进一步思考如何获得更高的排序性能，没错，就是 goroutine，将一个数据切分成两半，分别使用 `StdSortWithoutInterface` 排序，将排序后的结果进行一次归并排序，就可以得到最终的有序数组，这次我们测试的数组长度为 **10e5**

为了验证真正的`并行计算` 我们将分别测试 cpu 数量为 1, 2, 8 的情况
```shell
BenchmarkStdSort                               5         260696480 ns/op
BenchmarkStdSort-2                             5         246746560 ns/op
BenchmarkStdSort-8                             5         248532560 ns/op
BenchmarkStdSortWithoutInterface              10         124666470 ns/op
BenchmarkStdSortWithoutInterface-2            10         120676740 ns/op
BenchmarkStdSortWithoutInterface-8            10         126062650 ns/op
BenchmarkStdSortWithGoroutine                 20         125163280 ns/op
BenchmarkStdSortWithGoroutine-2               20          80835825 ns/op
BenchmarkStdSortWithGoroutine-8               20          81232625 ns/op
```

😎 WOW!!! cpu 数量为 1 时大家相差无几，cpu > 1 以后，goroutine 做到了真正的并行，利用多核进行计算，速度提升了 **1.5** 倍，比默认的 Sort 方法提升了 **4** 倍。诺，这就是算法的魅力。

### 📃 License
MIT [©chenjiandongx](http://github.com/chenjiandongx)
