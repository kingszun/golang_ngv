package main

import (
	"fmt"
	"time"
)

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		//1초 대기
		time.Sleep(time.Second)
		// 채널에 데이터를 기록
		// 버퍼에 기록을 했기 때문에 누가 가져가기 전까지는 블락되어 다음 작업을 수행 할 수 없음
		ch <- i

		fmt.Println("데이터 기록 성공")
		fmt.Println("보낸 데이터:", i)
	}
	// 채널 닫기 - 채널을 읽는 쪽에 시그널을 전송합니다.
	// 두번째 리턴 값이 false가 됩니다.
	close(ch)

}

func receive(ch chan int) {
	for {
		// 채널에서 데이터를 읽어오는데 첫번째 리턴값은 데이터
		// 두번째 리턴값은 데이터의 전송 여부
		// 데이터를 전송하면 첫번째 리턴에 데이터를 넘겨주고 뒤는 true가 됩니다.
		// 채널을 닫으면 데이터는 없고 뒤의 값이 false가 됩니다.
		// i, ok := <-ch
		// if ok == false {
		// 	break
		// }
		// fmt.Println("보낸 데이터:", i)
		if i, ok := <-ch; ok == false {
			break
		} else {
			fmt.Println("받은 데이터:", i)
		}

	}
}

func receiveRange(ch chan int) {
	for {
		for i := range ch {
			fmt.Println("받은 데이터:", i)
		}
	}
}

// goroutine의 실행 순서는 알 수 없음
// 실행 순서를 제어하려면 public subscribe 구조로 가야함
func main() {
	// 정수를 데이터로 전송하는 동기 채널을 생성
	myChan := make(chan int)

	go send(myChan)
	// go receive(myChan)
	receiveRange(myChan)

	fmt.Scanln() // 키보드 엔터가 올 때 까지 대기
}
