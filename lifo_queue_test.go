package go_collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLifoQueue(t *testing.T) {
	q := NewLifoQueue()

	var item interface{}
	var ok bool

	for i := 0; i < nums; i++ {
		q.Put(i)
	}
	for i := nums - 1; i >= 0; i-- {
		item, ok = q.Get()
		assert.Equal(t, i, item.(int))
		assert.Equal(t, ok, true)
	}

	item, ok = q.Get()
	assert.Equal(t, nil, item)
	assert.Equal(t, ok, false)
	
	item, ok = q.Get()
	assert.Equal(t, nil, item)
	assert.Equal(t, ok, false)
}
