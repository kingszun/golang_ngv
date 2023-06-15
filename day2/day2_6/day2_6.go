package main

import "fmt"

func pAttack() {
	fmt.Println("프로토스 어택")
}

func zAttack() {
	fmt.Println("저그 어택")
}

func main() {
	// 단축키를 만들려면 단축키가 2개 필요함
	pAttack()
	zAttack()

	// 매개변수가 없고 리턴이 없는 함수를 대입할 수 있는 변수가 생성
	// star라는 변수로 pAttack zAttack 모두 사용 가능
	// 동일한 원형의 함수를 여러 개 호출하고자 할 때 변수를 생성해서 하기도 합니다.
	var star func()
	star = pAttack
	star()
	star = zAttack
	star()
}
