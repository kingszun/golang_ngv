package benchmark

// 재귀로 구현한 피보나치 수열
func fibonacci1(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	} else {
		return fibonacci1(n-1) + fibonacci1(n-2)
	}
}

// 데이터 이동으로 구현한 피보나치 수열
func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	} else {
		one := 1
		two := 0
		rst := 0
		for i := 2; i <= n; i++ {
			rst = one + two
			two = one
			one = rst
		}
		return rst
	}
}
