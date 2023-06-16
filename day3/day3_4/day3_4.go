package main

import "fmt"

func main() {
	numbers := make(chan int, 5)
	for i := 0; i < 10; i++ {
		select {
		case numbers <- i * i:
			fmt.Println("작업 수행")
		default:
			fmt.Println("더 이상 저장 못함")
		}
	}
	for {
		select {
		case num := <-numbers:
			fmt.Println(num)
		default:
			fmt.Println("더이상 전송된 데이터가 없음")
			return
		}
	}

	fmt.Scanln()
}
