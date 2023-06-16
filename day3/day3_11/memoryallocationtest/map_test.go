package map_test

import "testing"

func test(m map[int]int) {
	//map에 데이터를 100000번 추가하는 함수
	for i := 0; i < 100000; i++ {
		m[i] = i
	}
}

func testSlice(sl []int) {
	for i := 0; i < 10000; i++ {
		sl = append(sl, i)
	}
}

func BenchmarkNoAllocationMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		test(m)
	}
}

func BenchmarkAllocationMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int, 440000)
		test(m)
	}
}
func BenchmarkNoAllocationSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make([]int, 1)
		testSlice(m)
	}
}
func BenchmarkAllocationSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make([]int, 1, 10000)
		testSlice(m)
	}
}
