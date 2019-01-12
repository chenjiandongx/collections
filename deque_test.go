package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeque(t *testing.T) {
	q := NewDeque()

	var item interface{}
	var ok bool

	for i := 0; i < nums; i++ {
		q.PutLeft(i)
	}
	assert.True(t, q.Qsize() == nums)

	for i := nums - 1; i >= 0; i-- {
		q.PutRight(i)
	}
	assert.True(t, q.Qsize() == nums*2)

	for i := 0; i < nums; i++ {
		item, ok = q.GetRight()
		assert.Equal(t, i, item)
		assert.Equal(t, ok, true)
	}
	for i := nums - 1; i >= 0; i-- {
		item, ok = q.GetLeft()
		assert.Equal(t, i, item)
		assert.Equal(t, ok, true)
	}
	assert.True(t, q.Qsize() == 0)

	item, ok = q.GetLeft()
	assert.Equal(t, nil, item)
	assert.Equal(t, ok, false)

	item, ok = q.GetRight()
	assert.Equal(t, nil, item)
	assert.Equal(t, ok, false)
}
