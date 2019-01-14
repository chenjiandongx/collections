package collections

import "sort"

type Counter struct {
	kv map[interface{}]int
}

type Item struct {
	k interface{}
	v int
}

func NewCounter() *Counter {
	return &Counter{make(map[interface{}]int)}
}

func (c *Counter) Add(keys ...interface{}) {
	for i := 0; i < len(keys); i++ {
		c.kv[keys[i]]++
	}
}

func (c *Counter) Get(key interface{}) int {
	return c.kv[key]
}

func (c *Counter) GetAll() []Item {
	return c.sortMap()
}

func (c *Counter) Top(n int) []Item {
	sortItems := c.sortMap()
	if n > c.Len() || n < 0 {
		n = c.Len()
	}
	return sortItems[:n]
}

func (c *Counter) Delete(key interface{}) bool {
	if _, ok := c.kv[key]; ok {
		delete(c.kv, key)
		return true
	}
	return false
}

func (c *Counter) Len() int {
	return len(c.kv)
}

func (c *Counter) sortMap() []Item {
	var items []Item
	for k, v := range c.kv {
		items = append(items, Item{k, v})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].v > items[j].v
	})
	return items
}
