package main

import (
	"fmt"
	"time"
)

// 시간과 관련된 패키지
func disp() {
	for i := 0; i < 10; i++ {
		//1초 대기
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

// 일반함수 호출은 하나의 함수 수행이 종료되어야만 다음 함수 호출이 가능
func disprun() {
	disp()
	disp()
}

// go 를 추가하면 고루틴으로 수행돼서 병행 수행이 될 수도 있고
// 하나의 작업 중 sleep이나 block(입출력) 상태가 되면 다른 goroutine을 수행할 수도 있습니다.
// goroutine은 백그라운드에서 작동하기 때문에 포그라운드가 있어야 작동 가능
func dispgoroutine() {
	go disp()
	go disp()
}

func send() {
	for i := 0; i < 10; i++ {
		//1초 대기
		time.Sleep(time.Second)
		fmt.Println(i)
	}

}

func receive() {

}

func main() {
	// disprun()

	dispgoroutine()

	time.Sleep(15 * time.Second) // Goroutine작업이 10초정도 걸리니깐 15처정도 대기
	fmt.Scanln()                 // 키보드 엔터가 올 때 까지 대기
}
