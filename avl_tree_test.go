package collections

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLTree(t *testing.T) {
	tree := NewAVLTree()
	for i := 0; i < maxNum; i++ {
		tree.Insert(i)
		tree.Insert(maxNum + i)
	}
	assert.Equal(t, len(tree.AllValues()), maxNum*2)
	assert.True(t, assertSort(tree.AllValues()))
	assert.Equal(t, tree.GetMaxValue(), 2*maxNum-1)
	assert.Equal(t, tree.GetMinValue(), 0)
	assert.True(t, tree.Search(50))
	assert.True(t, tree.Search(100))
	assert.False(t, tree.Search(-10))
	assert.False(t, tree.Delete(-10))
	assert.True(t, tree.Delete(10))
	assert.True(t, assertSort(tree.AllValues()))
	assert.Equal(t, len(tree.AllValues()), maxNum*2-1)
}

func TestAVLTreeRandom(t *testing.T) {
	tree := NewAVLTree()
	for i := 0; i < maxNum; i++ {
		tree.Insert(rand.Int())
	}
	assert.Equal(t, len(tree.AllValues()), maxNum)
	assert.True(t, assertSort(tree.AllValues()))
}

func genAVL(n int) *AVLTree {
	t := NewAVLTree()
	for i := 0; i < n; i++ {
		t.Insert(rand.Int())
	}
	return t
}

func BenchmarkAVLInsert10e1(b *testing.B) { genAVL(10e1) }
func BenchmarkAVLInsert10e2(b *testing.B) { genAVL(10e2) }
func BenchmarkAVLInsert10e3(b *testing.B) { genAVL(10e3) }
func BenchmarkAVLInsert10e4(b *testing.B) { genAVL(10e4) }
func BenchmarkAVLInsert10e5(b *testing.B) { genAVL(10e5) }

var at = genAVL(10e5)

func BenchmarkAVLSearch(b *testing.B) { at.Search(rand.Int()) }
func BenchmarkAVLDelete(b *testing.B) { at.Delete(rand.Int()) }
