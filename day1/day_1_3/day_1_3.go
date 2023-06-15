package main

import "fmt"

// 데이터의 저장 위치를 가리키는 이름을 생성
// {} 외부에서는 이런 형태로 생성
var age int = 53

func main() {
	//{} 내부에서 변수를 생성하는 방법
	var key int = 10
	heigth := 168
	var zerovalue int //기본값으로 생성실해

	fmt.Println(key, heigth, zerovalue)
	fmt.Printf("%d, %d, %d \n", key, heigth, zerovalue)
}
