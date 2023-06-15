package main

import "fmt"

type super struct {
	num  int
	name string
}

type sub struct {
	super
	score int
}

// super의 메서드
func (s *super) display() {
	fmt.Println("super의 메서드")
}

// super의 메서드
func (s *sub) display() {
	s.super.display() //super 것을 호출하고자 하면 하위구조체이름.상위구조체이름.메서드() 로 호출
	fmt.Println("sub의 메서드")
}
func main() {
	s := new(sub)
	s.num = 1
	s.name = "adam"
	s.score = 99
	fmt.Println(s)
	// sub에서 display를 찾고 있으면 sub의 display를 호출
	// sub에서 display를 찾고 없으면 super의 display를 호출
	s.super.display()
}
