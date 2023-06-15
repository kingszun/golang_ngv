package main

import "fmt"

func f() {
	// Hello World 10번 출력
	for i := 0; i < 10; i++ {
		fmt.Println("Hello World")
	}
}

// 매개변수가 있으면 매개변수를 이용해 다양한 작업을 수행할 수 있습니다.
// 매개변수는 호출할 때 복사해서 사용하게 됩니다.
func fn(n int) {
	// Hello World n번 출력
	for i := 0; i < n; i++ {
		fmt.Println("Hello World")
	}
}

// return 이 있는 경우 리턴되는 데이터를 이용해서 다음 작업을 수행할 수 있습니다.
func add(n1 int, n2 int) {
	result := n1 + n2
	fmt.Println(result)
}
func radd(n1 int, n2 int) int {
	result := n1 + n2
	fmt.Println(result)
	return result
}

// return : 함수의 수행이 종료되고 호출하는 곳으로 돌아가는 명령어
// 리턴을 할 때 데이터가 없을 수도 있지만 데이터를 가지고 갈 수 도 있습니다.

// 2개의 정수를 받아 덧셈을 한 결과를 리턴해주는 함수
func radd1(n1 int, n2 int) int {
	return n1 + n2
}

// go는 리턴 되는 데이터에 이름을 붙일 수 있습니다.
// 리턴 타입 앞에 이름을 만들면 내부에 변수가 자동으로 생성되고 자동으로 리턴됩니다.
// 함수의 시그니쳐만 보고 함수의 입/출력을 예측 가능하기 때문에 이런 방식으로 쓰는걸 추천
func radd2(n1 int, n2 int) (result int) {
	// result 는 위에서 자동 생성된다.
	// 리턴되는 자료형 앞에 이름을 만들면 이 데이터는 자동으로 리턴
	result = n1 + n2
	return // 함수선언시 만든 이름은 자동으로 리턴됩니다.
}

// 2개가 리턴되지만 이런식의 코드는 공부할때만 짜고 실제에선 거의 짜지 않는다.
// 함수, 구조체, 클래스는 1개의 역할만 하는게 좋다.
// 그런데 이렇게 되면 2개의 역할을 하게돼서 좋지 않음
func calc(n1 int, n2 int) (a int, b int) {
	a = n1 + n2
	b = n1 - n2
	return
}
func main() {
	//Hello World 를 10번 출력해보자
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	// 이게 가장 간단하고 빠름
	// 그런데 이렇게 만들면 유지보수가 너무 힘듬

	// 함수 사용하는 이유
	// 모듈화
	// 중복된거 방지
	// Hello World 10번찍는걸 2번 할경우
	// 이렇게 되면 Hello World 를 변경할때 따로 2번 고쳐야함
	for i := 0; i < 10; i++ {
		fmt.Println("Hello World")
	}
	fmt.Println("Hi")
	for i := 0; i < 10; i++ {
		fmt.Println("Hello World")
	}

	//함수 호출 - 별도의 메모리 공간을 할당받아서 수행
	// 해당함수는 정해진 수행밖에 못하지만 매개변수를 통해 다양한 동작 가능
	f()
	fmt.Println("Hi")
	f()

	// Hello World를 5번, 7번 출력
	// 매개변수가 있으면 매개변수를 이용해 다양한 작업을 수행할 수 있습니다.
	fn(5)
	fn(7)

	// 매개변수는 호출할 때 복사해서 사용하게 됩니다.
	var x int = 3
	// x가 대입되는 것이 아니고 3이 대입됩니다.
	// 함수실행도중 x가 바뀐다 하더라도 함수실행에는 변동이 없습니다.
	fn(x)

	// 리턴이 없는 함수를 만들경우 함수는 실행되고 끝임.
	// 해당 데이터를 재활용할 방법이 없음
	add(10, 20)

	// 리턴이 있는 함수를 만들면 작업을 이어서 할 수 있음
	radd(radd(10, 20), 30)
	fmt.Println(radd(radd(10, 20), 50))

	// 변수 선언
	var a1 int = 10
	var a2 int = 20
	// 함수 호출
	result := radd1(a1, a2)
	fmt.Println(result)
	r1, r2 := calc(a1, a2)
	fmt.Println(r1, r2)
	var r3 int = 0
	r3, _ = calc(a1, a2)
	_, r4 := calc(a1, a2)
	fmt.Println(r3, r4)

	// err이 있는 함수는 아래와 같이 쓰는게 좋다.
	if n, err := fmt.Println("Hello", "World"); err == nil {
		//성공했을 때
		fmt.Println(n)
	} else {
		//실패했을 때
		fmt.Println(err)
	}
}
