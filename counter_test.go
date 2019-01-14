package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyCounter(t *testing.T) {
	c := NewCounter()
	assert.Equal(t, c.Len(), 0)
	assert.Equal(t, len(c.GetAll()), 0)
	assert.Equal(t, c.Get("anything"), 0)
}

func TestCounter(t *testing.T) {
	c := NewCounter()
	c.Add("a", "b", "c", "d", "a", "c")
	assert.Equal(t, c.Top(2), []Item{{"a", 2}, {"c", 2}})
	assert.Equal(t, c.Get("A"), 0)
	assert.Equal(t, c.Get("a"), 2)
	assert.Equal(t, c.Get("b"), 1)
	assert.Equal(t, c.Len(), 4)
	assert.Equal(t, len(c.Top(10)), c.Len())
	assert.Equal(t, len(c.Top(-10)), c.Len())
	c.Delete("a")
	assert.Equal(t, c.Get("a"), 0)
}
