package main

import "fmt"

func main() {
	//1부터 3까지 순서대로 출력
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
	// i<=3이 false가 될때까지 진행
	// i<=3이 true라서 진행이 아니라 false가 아니라서 진행
	// i=4일때 멈춘다
	// 보통 내 생각보다 1번 더 실행됨
	var j int
	for i := 1; i < 4; i, j = i+1, j+1 {
		fmt.Println(i)
		fmt.Println(j)
	}

	for j := 1; j <= 10; j++ {
		// break는 반복문 중단
		// 1과 2를 출력하고 3이 되면 반복문이 중단
		// continue는 아래 문장을 수행하지 않음
		if j%3 == 0 {
			break
		}
		fmt.Println(j)
	}
	for j := 1; j <= 10; j++ {
		// continue는 아래 문장을 수행하지 않음
		if j%3 == 0 {
			continue
		}
		fmt.Println(j)
	}
	return //함수의 수행을 종료하고 호출한 곳으로 돌아가는 명령어
	//main 에서는 프로그램 종료의 의미를 갖는다
	//main 에서 0을 반환하면 정상실행 0이아닌 정수를 리턴하면 에러 반환
	fmt.Println("Hello")
}
