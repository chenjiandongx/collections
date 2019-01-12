package collections

type linkedList struct {
	next, prev *linkedList
	key, value interface{}
}

type OrderedMap struct {
	head, tail, current *linkedList
	len                 int
	items               map[interface{}]*linkedList
}

func NewOrderedMap() *OrderedMap {
	headNode := &linkedList{}
	tailNode := &linkedList{}
	headNode.next, tailNode.prev = tailNode, headNode
	return &OrderedMap{
		head:    headNode,
		tail:    tailNode,
		current: headNode,
		len:     0,
		items:   make(map[interface{}]*linkedList),
	}
}

func (om *OrderedMap) Set(key, value interface{}) {
	if _, ok := om.items[key]; ok {
		om.items[key].value = value
		return
	}
	newNode := &linkedList{prev: om.tail, next: nil, key: key, value: value}
	om.items[key] = newNode
	om.tail.next = newNode
	om.tail = newNode
	om.len ++
}

func (om *OrderedMap) Get(key interface{}) (interface{}, bool) {
	if v, ok := om.items[key]; ok {
		return v.value, ok
	}
	return nil, false
}

func (om *OrderedMap) Delete(key interface{}) bool {
	item, ok := om.items[key]
	if !ok {
		return ok
	}
	if item.next == nil {
		om.tail.prev, om.tail.next = om.tail, nil
	} else {
		item.prev.next, item.next.prev = item.next, item.prev
	}
	delete(om.items, key)
	item = nil
	om.len --
	return true
}

func (om *OrderedMap) Iter() (interface{}, interface{}, bool) {
	c := om.current.next
	if c.next != nil {
		om.current = om.current.next
		return c.next.key, c.next.value, true
	}
	return nil, nil, false
}

func (om *OrderedMap) BackToHead() {
	om.current = om.head
}

func (om *OrderedMap) Len() int {
	return om.len
}
