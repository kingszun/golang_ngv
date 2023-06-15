package main

import "fmt"

func main() {
	ar := [3]string{"사이버가수아담", "강진축구", "사이버컵"}
	//강진 축구 출력
	fmt.Println(ar[1])
	//배열의 데이터 개수
	fmt.Println(len(ar))
	// 순회 - 하나씩 모두 접근
	for _, value := range ar {
		fmt.Println(value)
	}
	for index := range ar {
		fmt.Println(index)
	}
	for index, _ := range ar {
		fmt.Println(index)
	}

	br := ar //배열은 대입하면 데이터 전체를 복사해서 줍니다.
	fmt.Println(ar)
	fmt.Println(br)
	br[2] = "프리스톤테일" //배열이므로 ar의 값은 변경되지 않습니다.
	fmt.Println(ar)
	fmt.Println(br)

	sl := ar[0:3] //배열 이용해서 슬라이스 생성
	fmt.Println(sl)

	sl1 := sl
	fmt.Println(sl1)
	sl1[0] = "아담"
	fmt.Println("sl", sl)
	fmt.Println("sl1", sl1)

	fmt.Println(len(sl), cap(sl))
	//데이터를 추가할 때 빈 공간이 없으면 현재 공간의 2배를 할당하고 데이터를 추가
	sl = append(sl, "프리스톤테일")
	fmt.Println(len(sl), cap(sl))

	x := make([]string, 3, 5) //빈 문자열 3개를 가지고 용량은 5인 슬라이스 생성
	// 데이터 2개 추가할때까지는 용량을 늘리지 않음
	// 잘못만들면 메모리누수 생기고 느려짐 용량을 잘 정하자
	fmt.Println(len(x), cap(x))
	x = append(x, "아담소프트")
	fmt.Println(len(x), cap(x))
	x = append(x, "넥슨")
	fmt.Println(len(x), cap(x))
	x = append(x, "카카오엔터테인먼트")
	fmt.Println(len(x), cap(x))

	y := []string{}
	y = append(y, "a")
	fmt.Println(len(y), cap(y))
	y = append(y, "a")
	fmt.Println(len(y), cap(y))
	y = append(y, "a")
	fmt.Println(len(y), cap(y))
	y = append(y, "a")
	fmt.Println(len(y), cap(y))
	y = append(y, "a")
	fmt.Println(len(y), cap(y))
	y = append(y, "a")
	fmt.Println(len(y), cap(y))
	y = append(y, "a")
	fmt.Println(len(y), cap(y))
	y = append(y, "a")
	fmt.Println(len(y), cap(y))

	sl2 := ar[0:2]
	sl2[0] = "아담"

	fmt.Println(ar)
	//기존 len과 cap가 동일한상태에서 데이터를 추가하면 기존 배열과의 연결이 끊어집니다.
	sl2 = append(sl2, "프리스톤테일", "넥슨")
	sl2[0] = "이상철"
	fmt.Println(ar)
	fmt.Println(sl2)

}
