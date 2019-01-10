package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const nums = 1000

func TestQueue(t *testing.T) {
	q := NewQueue()

	var item interface{}
	var ok bool

	for i := 0; i < nums; i++ {
		q.Put(i)
	}
	for i := 0; i < nums; i++ {
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
