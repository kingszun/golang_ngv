package main

import "fmt"

type data struct {
	num  int
	num2 int8
	num3 float32
}

func main() {
	var value int = 0x0001000100010001
	var pointer *int = &value // 다른 데이터의 참조를 가지고 포인터 생성
	fmt.Println(pointer)      // 포인터가 가리키는 참조
	fmt.Println(*pointer)     // 참조에 저장된 데이터
	fmt.Println(&pointer)     //포인터 자체의 참조

	pointer = new(int) //int형 메모리 한개를 할당받아서 참조를 대입
	*pointer = 100
	fmt.Println(pointer)
	fmt.Println(*pointer)
	fmt.Println(&pointer)
	*pointer = 10
	fmt.Println(pointer)
	fmt.Println(*pointer)
	fmt.Println(&pointer)

	var str string = "안녕하세요"
	var p *string = &str
	fmt.Println(p)

	// 포인터를 출력했을 때 가끔은 &{데이터} 형태로 출력됩니다.
	var one data
	var two *data = &one
	fmt.Println(two) //구조체에는 데이터가 여러개 쌓여있기 때문에 참조보다는 데이터를 알려줌
	fmt.Printf("%p \n", two)
}
