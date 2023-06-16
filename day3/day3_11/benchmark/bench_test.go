package benchmark

import "testing"

// b.N으로 작성하면 go가 1,2,5,10,50,... 이런식으로 숫자를 늘려가면서 함수를 호출해서 수행시간의 평균을 구하게 됩니다.
func Benchmark_Fibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}

func Benchmark_Fibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}
