package main

import "fmt"

// 프로그램을 실행할 때 만들어서 메모리 할당을 받고 프로그램이 종료되면 소멸
// 이곳에 만들면 전역으로 선언되기 때문에 해제가 불가능하다.
// 여러곳에서 호출되면 이게 좋음
func display() {
	// display가 불려온 상태에서만 있기 때문에 main에서 바꿀 수 없음
	// closure를 이용해서 main에서 바꿀 수 있음.
	var a int
	fmt.Println("일반적인 함수", a)
}

// outer 함수 안에서 함수를 만들어서 리턴
// 함수가 함수를 리턴
// 리턴하는 함수가 외부 함수의 데이터를 수정한다 : closure
// outer를 호출하면 함수가 리턴된다
func outer() func() {
	var a int = 0
	// 내부함수에서는 외부함수의 데이터를 사용 할 수 있음
	// closure를 이용하면 함수 외부에서 함수 내의 데이터를 변경할 수 있습니다.
	// 전역 변수를 사용하지 않고 데이터를 공유하는 방식을 고민하자!
	return func() {
		// 내부함수
		a = a + 1
		fmt.Println(a)
	}
}
func main() {

	display()

	// 익명 함수를 만들어서 호출
	// 여기서 선언되면 메모리 해제가 가능하다.
	// 호출되는 시점에 만들어지고 감싸고 있는 함수가 종료되면 소멸
	// 한번쓰고 마는 애들은 이게 좋음
	func() {
		fmt.Println("익명 함수")
	}()

	f := outer()
	f()
	f()
	d := outer()
	d()
}
