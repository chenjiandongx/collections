package go_collections

import (
	"math/rand"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	q := NewPriorityQueue()

	for i := 0; i < nums; i++ {
		r := rand.Int()
		q.Put(&PqNode{Value: string(r), Priority: rand.Int()})
	}

	for i := 0; i < nums/2; i++ {
		item1, ok := q.Get()
		item2, ok := q.Get()
		assert.True(t, item1.(*PqNode).Priority > item2.(*PqNode).Priority)
		assert.Equal(t, ok, true)
	}
	var item interface{}
	var ok bool

	item, ok = q.Get()
	assert.Equal(t, nil, item)
	assert.Equal(t, ok, false)

	item, ok = q.Get()
	assert.Equal(t, nil, item)
	assert.Equal(t, ok, false)
}
