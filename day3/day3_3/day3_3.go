package main

import (
	"fmt"
	"time"
)

func main() {
	// 문자열과 정수를 전송할 수 있는 채널을 생성
	ch1 := make(chan string)
	ch2 := make(chan int)

	// 익명 함수를 만들자 마자 실행
	// 1초마다 "안녕하세요"를 ch1에 계속 기록
	go func() {
		for {
			ch1 <- "안녕하세요"
			time.Sleep(time.Second)
		}
	}()
	// 0.5초마다 정수를 1씩 증가시키며 ch2에 기록
	go func() {
		i := 0
		for {
			i++
			ch2 <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()
	// 2개의 채널로부터 데이터 읽기
	// go func() {
	// 	for {
	// 		i := <-ch1
	// 		fmt.Println(i)
	// 		j := <-ch2
	// 		fmt.Println(j)
	// 	}
	// }()
	// 독립적으로 작동하게 할 때
	// go func() {
	// 	for {
	// 		// select를 사용하면select에서 블럭되고
	// 		// 체널에 데이터가 오면 그 구문을 수행하고 다시 select로 돌아오게 만들 수 있습니다.
	// 		select {
	// 		case i := <-ch1:
	// 			fmt.Println(i)
	// 		case j := <-ch2:
	// 			fmt.Println(j)
	// 		}
	// 	}
	// }()

	go func() {
		for {
			select {
			case i := <-ch1:
				fmt.Println(i)
			case j := <-ch2:
				fmt.Println(j)
			//다른 채널이 데이터를 보내기 전에 이 함수가 호출되도록 만들면 다른 채널에 데이터를 전송하더라도 읽을 수 없습니다.
			case <-time.After(300 * time.Microsecond):
				fmt.Println("이 시간동안 데이터가 오지 않으면 패스")
			}
		}
	}()

	fmt.Scanln()
}
