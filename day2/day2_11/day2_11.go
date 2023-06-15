package main

import "fmt"

// 인터페이스는 함수의 원형만 소유
type protocol interface {
	display()
}

// data 구조체는 protocol 인터페이스를 implements 했다라고 합니다.
// protocol이라는 프로토콜을 conform 했다고도 합니다.
type data struct {
}

// data 구조체에 protocol의 메서드를 구현
func (d data) display() {
	fmt.Println("data의 display")
}

type chat struct {
}

func (c chat) display() {
	fmt.Println("chat의 display")
}

// 빈 인터페이스 - 구현알 내용이 없는 인터페이스
// 이 인터페이스 변수에는 모든 데이터를 저장 할 수 있습니다.
// interface{} - 이게 빈 인터페이스라는걸 기억해야함
type bean interface {
}

func main() {
	//인터페이스 변수는 참조형
	// 인터페이스는 내용이 없어서 메모리 할당을 받을 수 없음
	// var p protocol = new(protocol)
	var p protocol
	fmt.Println(p)

	//data 구조체가 protocol을 implements 했기 때문에 대입이 가능
	p = data{}
	p.display()

	//동일한 코드가 대입된 인스턴스에 따라 다른 메서드를 호출 - 다형성
	// 자료구조처럼 일을 하는 것 자체는 같지만 내부 구현이 다른경우
	// 이 방식을 이용하면 사용자가 편리합니다.
	p = chat{}
	p.display()

	var b bean
	b = 10
	fmt.Println(b)
	// 10은 int인데b는 현재 bean 인터페이스
	fmt.Println(b.(int) + 10)
	b = "Hello World"
	fmt.Println(b)
	fmt.Println(b.(string))
}
