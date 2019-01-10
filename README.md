# go-queue

[![GoDoc](https://godoc.org/github.com/chenjiandongx/go-queue?status.svg)](https://godoc.org/github.com/chenjiandongx/go-queue)
> Golang 实现的线程安全队列

### Install
```bash
$ go get github.com/chenjiandongx/go-queue
```

### Import
```go
import Q "github.com/chenjiandongx/go-queue"
```

#### Method

所有队列都实现了以下方法
```go
Get()(interface{}, bool)
Put(v interface{})
Qsize() int
IsEmpty() bool
```

#### Queue
> 先进先出队列

```go
var nums = 1000

q := Q.NewQueue()
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

#### LifoQueue
> 后进先出队列

```go
var nums = 1000

q := Q.NewLifoQueue()
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


#### PriorityQueue
> 优先队列

```go
var nums = 1000

q := Q.NewPriorityQueue()

for i := 0; i < nums; i++ {
    r := rand.Int()
    q.Put(&Q.PqNode{Value: string(r), Priority: rand.Int()})
}

for i := 0; i < nums/2; i++ {
    item1, _ := q.Get()
    item2, _ := q.Get()
    fmt.Println(item1.(*Q.PqNode).Priority > item2.(*Q.PqNode).Priority)
}

fmt.Println(q.IsEmpty())
fmt.Println(q.Qsize())
```

### License
MIT [©chenjiandongx](http://github.com/chenjiandongx)
