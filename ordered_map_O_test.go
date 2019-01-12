package collections

import (
	"runtime"
	"sort"
	"testing"
)

type tester interface {
	has(int) bool
}

type mapTester struct {
	m map[int]bool
}

func (m *mapTester) has(x int) bool {
	_, ok := m.m[x]
	return ok
}

type sliceTester struct {
	s []int
}

func (s *sliceTester) has(x int) bool {
	// sort.Search 使用的是二分查找
	i := sort.Search(len(s.s), func(i int) bool { return s.s[i] >= x })
	// 未找到时 i 为 -1
	return i < len(s.s) && s.s[i] == x
}

func genMap(size int) tester {
	m := map[int]bool{}
	for i := 0; i < size; i++ {
		m[i] = true
	}
	return &mapTester{m: m}
}

func genSlice(size int) tester {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = i
	}
	return &sliceTester{s: s}
}

func benchmark(r tester, b *testing.B) {
	runtime.GC()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.has(i)
	}
}

func BenchmarkMap10e1(b *testing.B)   { benchmark(genMap(10e1), b) }
func BenchmarkSlice10e1(b *testing.B) { benchmark(genSlice(10e1), b) }
func BenchmarkMap10e2(b *testing.B)   { benchmark(genMap(10e2), b) }
func BenchmarkSlice10e2(b *testing.B) { benchmark(genSlice(10e2), b) }
func BenchmarkMap10e3(b *testing.B)   { benchmark(genMap(10e3), b) }
func BenchmarkSlice10e3(b *testing.B) { benchmark(genSlice(10e3), b) }
func BenchmarkMap10e4(b *testing.B)   { benchmark(genMap(10e4), b) }
func BenchmarkSlice10e4(b *testing.B) { benchmark(genSlice(10e4), b) }
func BenchmarkMap10e5(b *testing.B)   { benchmark(genMap(10e5), b) }
func BenchmarkSlice10e5(b *testing.B) { benchmark(genSlice(10e5), b) }
func BenchmarkMap10e6(b *testing.B)   { benchmark(genMap(10e6), b) }
func BenchmarkSlice10e6(b *testing.B) { benchmark(genSlice(10e6), b) }
func BenchmarkMap10e7(b *testing.B)   { benchmark(genMap(10e7), b) }
func BenchmarkSlice10e7(b *testing.B) { benchmark(genSlice(10e7), b) }
