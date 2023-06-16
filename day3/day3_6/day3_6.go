package main

import (
	"fmt"
	"runtime"
	"sync"
)

// "runtime" 현재 go의 환경이나 스레드 개수, CPU를 앙됴할 때 사용

func main() {
	var data = []int{} // 정수형 슬라이스

	var mutex = new(sync.Mutex)

	go func() {
		// 슬라이스에 1을 100개 추가
		for i := 0; i < 1000; i++ {

			mutex.Lock() // 아래 코드 영역이 실행되는 동안 데이터를 다른 곳에서 수정할 수 없음
			data = append(data, 1)
			mutex.Unlock() // 다른곳에서 데이터를 수정할 수 있음
			//CPU양도
			runtime.Gosched()
		}
	}()
	go func() {
		// 슬라이스에 1을 100개 추가
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()
			//CPU양도
			runtime.Gosched()
		}
	}()

	fmt.Scanln()
	fmt.Println(len(data))
}
