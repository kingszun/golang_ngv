package main

import "fmt"

func callbyvalue(n int) {
	n = n + 10
	fmt.Println("함수 내부 n:", n)
}

func callbyreference(p *int) {
	*p = *p + 10
	fmt.Println("함수 내부 *p:", *p)
}

// 이렇게 만들면 이상한 행동을 한거다
// 굳이 바로 출력하면 되는걸 한단계 더 거쳐갈 필요가 없음
func callbyreference2(p *int) {
	fmt.Println("함수 내부 *p:", *p)
}
func main() {
	var a int = 10
	// 함수에 값을 전달했으므로 a의 값은 그대로
	callbyvalue(a)
	fmt.Println("a:", a)

	// 함수에 reference를 전달했으므로 a의 값이 변경되었을 수 도 있습니다.
	callbyreference(&a)
	fmt.Println("a:", a)

	// 함수 이름은 함수가 있는 곳의 참조
	fmt.Println(callbyvalue) //함수가 저장된 정적 메모리 공간을 출력해줌
}
