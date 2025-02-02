package main

import "fmt"

func main() {
	//상수 선언
	const NAME = "아담"

	//나열형 상수 선언
	//직접 데이터 설정
	const (
		MIN    = 0
		NORMAL = 1
		MAX    = 2
	)
	fmt.Println(MAX)

	//일련번호 형태로 설정
	const (
		ZERO = iota //0
		ONE         //1
		TWO         //2
	)
	fmt.Println(TWO)

	const (
		LBUTTON = 1 << iota //1
		RBUTTON = 1 << iota //2
		SHIFT   = 1 << iota //4
	)
	fmt.Println(SHIFT)

	//연산을 수행해서 일정한 패턴 적용
	const (
		T10 = iota * 10 //0
		T20 = iota * 10 //10
		T30 = iota * 10 //20
	)
	fmt.Println(T30)

	//특정 값 생략
	const (
		bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0
		bit1, mask1                          // bit1 == 2, mask1 == 1
		_, _                                 // iota == 2를 생략
		bit3, mask3                          // bit3 == 8, mask3 == 7
	)
	fmt.Println(bit0)
	fmt.Println(mask0)
	fmt.Println(bit1)
	fmt.Println(mask1)
	fmt.Println(bit3)
	fmt.Println(mask3)
}
