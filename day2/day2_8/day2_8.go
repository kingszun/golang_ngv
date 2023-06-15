package main

import (
	"fmt"
)

func disp() {
	fmt.Println("defer")
}

func main() {
	fmt.Println("시작")
	// 강제로 예외를 발생
	panic("패닉 발생")

	// disp가 먼저 있지만 defer가 있으므로 수행이 맨 뒤로 밀리고 무조건 호출됨
	defer disp()
	// panic("패닉 발생")
	fmt.Println("종료?")
}
