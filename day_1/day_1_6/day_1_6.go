package main

import (
	"fmt"
)

func main() {
	a := 10 //10이라는 데이터를 가리키는 a 변수 생성
	a = 20  // ark 20이라는 데이터를 가리키도록 변경

	//a 는 intdlrh 20.8은 실수라서 float64입니다.
	// 자료형이 다르기 때문에 연산이 불가능합니다.
	// a 를 float64로 변경하거나 20.8을 int로 변경해서 수행해야합니다.
	result := float64(a) + 20.8
	fmt.Println(a)
	fmt.Println(result)
	a++
	fmt.Println(a)

	// 1부터 100까지의 숫자 중 3의 배수이고 4의 배수인 데이터를 출력
	// 아래 두 loop는 동일한 값이 출력됨
	// 그러나 실행되는 숫자는 i%4를 먼저 사용하는것이 리소스를 더 적게 사용함
	var i int = 1

	for i = 1; i < 100; i++ {
		if i%3 == 0 && i%4 == 0 {
			fmt.Println(i)
		}
	}
	for i := 1; i < 100; i++ {
		if i%4 == 0 && i%3 == 0 {
			fmt.Println(i)
		}
	}

}
