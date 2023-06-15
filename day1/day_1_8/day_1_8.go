package main

import "fmt"

func main() {
	var menu int = 1

	switch menu {
	case 1, 5: // 여러개의 값 설정 가능
		fmt.Println("한식")
		fallthrough // 아래도 실행시키기 위해선 fallthrough 사용해야함
	case 2:
		fmt.Println("양식")
	case 3:
		fmt.Println("분식")
	default:
		fmt.Println("굶기")
	}

	//switch에 true를 기재하거나 표현식을 생략하면 case에 bool 표현식 사용 가능
	switch true {
	case menu >= 4:
		fmt.Println("1")
	case menu < 4:
		fmt.Println("2")

	}
}
