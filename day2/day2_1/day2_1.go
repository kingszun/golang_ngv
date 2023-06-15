package main

import "fmt"

func main() {
	//초기 데터를 가지고 슬라이스 생성
	ar := []string{"월요일", "화요일", "수요일", "목요일", "금요일"}
	fmt.Println(ar)

	// //기존 슬라이스에서 앞의 2개의 데이터만 선택
	// fmt.Println(ar[0:2]) // [:2] 도 동일한 결과 default 0

	// //목요일과 금요일 선택하기
	// fmt.Println(ar[3:])

	xr := ar[0:2]
	// yr := []string{"토요일", "일요일"}
	// xr = append(xr, "아무 요일", "두번째 아무 요일")
	// xr = append(xr, yr...) // 가변인자에 슬라이스를 추가할때는 "..."을 기재해줘야 합니다.
	xr = append(xr, ar[3:5]...)
	fmt.Println(xr)

}
