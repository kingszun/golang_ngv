package coverage

import "fmt"

// 문법이나 동작 자체로는 아무런 문제가 없음
func f1() {
	if true {
		fmt.Println("무조건 수행")
	} else {
		fmt.Println("절대 도달할 수 없는 코드")
	}
}

func f2(n int) int {
	if n >= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return f2(n-1) + f2(n-2)
	}
}
