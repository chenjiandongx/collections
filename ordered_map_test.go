package collections

import (
	"testing"

	"github.com/cevaris/ordered_map"
)

const maxNum = 100

func BenchmarkCollectionsSet(b *testing.B) {
	om := NewOrderedMap()
	for i := 0; i < b.N; i++ {
		om.Set(i, i)
	}
}

func BenchmarkCevarisSet(b *testing.B) {
	om := ordered_map.NewOrderedMap()
	for i := 0; i < b.N; i++ {
		om.Set(i, i)
	}
}

func BenchmarkCollectionsGet(b *testing.B) {
	om := NewOrderedMap()
	for i := 0; i < b.N; i++ {
		om.Set(i, i)
	}
	for i := 0; i < b.N; i++ {
		om.Get(i)
	}
}

func BenchmarkCevarisGet(b *testing.B) {
	om := ordered_map.NewOrderedMap()
	for i := 0; i < b.N; i++ {
		om.Set(i, i)
	}
	for i := 0; i < b.N; i++ {
		om.Get(i)
	}
}

func BenchmarkCollectionsIter(b *testing.B) {
	om := NewOrderedMap()
	for i := 0; i < b.N; i++ {
		om.Set(i, i)
	}
	for k, v, ok := om.Iter(); ok; k, v, ok = om.Iter() {
		_, _ = k, v
	}
}

func BenchmarkCevarisIter(b *testing.B) {
	om := ordered_map.NewOrderedMap()
	for i := 0; i < b.N; i++ {
		om.Set(i, i)
	}
	iter := om.IterFunc()
	for kv, ok := iter(); ok; kv, ok = iter() {
		_, _ = kv.Key, kv.Value
	}
}

func TestSet(t *testing.T) {
	om := NewOrderedMap()
	for i := 0; i < maxNum; i++ {
		om.Set(i, i+1)
	}
	if om.len != maxNum {
		t.Error()
	}
}

func TestGet(t *testing.T) {
	om := NewOrderedMap()
	for i := 0; i < maxNum; i++ {
		om.Set(i, i+1)
	}
	for i := 0; i < maxNum; i++ {
		v, ok := om.Get(i)
		if !ok || v.(int) != i+1 {
			t.Error()
		}
	}
}

func TestGUpdate(t *testing.T) {
	om := NewOrderedMap()
	for i := 0; i < maxNum; i++ {
		om.Set(i, i+1)
	}
	for i := maxNum - 1; i >= 0; i-- {
		om.Set(i, i-1)
	}
	for i := 0; i < maxNum; i++ {
		v, ok := om.Get(i)
		if !ok || v.(int) != i-1 {
			t.Error()
		}
	}
}

func TestIter(t *testing.T) {
	om := NewOrderedMap()
	for i := 0; i < maxNum; i++ {
		om.Set(i, i+1)
	}
	index := 0
	for k, v, ok := om.Iter(); ok; k, v, ok = om.Iter() {
		if k.(int) != index || v.(int) != index+1 || !ok {
			t.Error()
		}
		index += 1
	}

	index = 0
	om.BackToHead()
	for k, v, ok := om.Iter(); ok; k, v, ok = om.Iter() {
		if k.(int) != index || v.(int) != index+1 || !ok {
			t.Error()
		}
		index += 1
	}
}
