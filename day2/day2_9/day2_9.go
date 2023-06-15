package main

import "fmt"

// 번호, 이름, 점수를 저장하기 위한 자료형을 선언
type person struct {
	num   int
	name  string
	score int
}

// person 구조체의 인스턴스의 참조를 리턴해주는 생성자와 유사한 함수
func NewPerson(num int, name string, score int) *person {
	return &person{num, name, score}
}

// person 의 메서드
func (p *person) getName() string {
	return p.name
}
func (p *person) setName(name string) {
	p.name = name
}
func main() {
	// 구조체 인스턴스 생성
	data1 := person{} // 기본값을 가지고 생성
	data2 := person{1, "개똥이", 100}

	// 구조체 인스턴스 참조 생성
	pData1 := new(person)
	pData2 := &person{2, "soccer", 99}

	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println(pData1)
	fmt.Println(pData2)

	//구조체 인스턴스나 참조를 이용해 필드 접근
	//go에서는 구조체 필드 접근은 무조건 "."을 사용함
	fmt.Println(data2.name)
	fmt.Println(pData2.score)

	cons := NewPerson(100, "itstudy", 89)
	fmt.Println(cons)

	// cons의 메서드 호출
	// set으로 시작하는건 보통 필드의 값을 수정
	// get으로 시작하는건 보통 필드의 값을 리턴
	fmt.Println(cons.getName())
	cons.setName("itggangpae")
	fmt.Println(cons.getName())
}
