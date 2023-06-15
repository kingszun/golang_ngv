package main

import "fmt"

func main() {
	fmt.Println("1")
	//1출력한 다음 2번을 수행하지 않고 3으로 이동하고 싶음
	goto JUMP //JUMP로 이동
UP:
	fmt.Println("2")

	//레이블 부여
JUMP:
	fmt.Println("3")
	goto UP
}
