package main

import (
	"fmt"
	"sync"
)

func main() {
	// WaitGroup 구조체 생성
	wg := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		wg.Add(1) //수행중인 goroutine이 1개 추가되었다고 표시
		go func(n int) {
			fmt.Println(n)
			// Goroutine의 수행이 완료되었다고 알려서 수행중인 goroutine의 개수를 1 감소시킵니다.
			// Add의 개수보다 Done이 많아지면 panic에 걸립니다.
			wg.Done()
		}(i)
	}
	//WaitGroup 이 관리하는 카운터가 0이 될 때 까지 대기
	// Goroutine과는 아무런 상관이 없음
	wg.Wait()

	fmt.Println("프로그램 종료")
}
