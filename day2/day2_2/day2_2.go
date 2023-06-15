package main

import "fmt"

func main() {
	//초기 데이터를 가지고 슬라이스 생성
	ar := []string{"월요일", "화요일", "수요일"}

	// 슬라이스는 reference type이므로 참조를 복사
	br := ar

	//복사본의 변경이 원본에 영향을 줌
	br[2] = "목요일"
	fmt.Println(ar)

	cr := make([]string, 3)

	//데이터를 새로운 공간에 복제해서 대입
	copy(cr, ar)
	fmt.Println(ar)
	//복사본의 변경이 원본에 영향을 주지 않습니다.
	cr[2] = "토요일"
	fmt.Println("ar:", ar)
	fmt.Println("cr:", cr)
	// cr은 ar의 데이터를 복제해서 만들어진 것으로 영향을 받지 않음
	ar[0] = "목요일"
	fmt.Println("ar:", ar)
	fmt.Println("cr:", cr)
}
