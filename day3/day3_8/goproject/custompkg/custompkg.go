package custompkg

import "fmt"

// 이 함수를 다른 패키지에서 사용하려면 함수 이름의 첫글자가 대문자이어야 합니다.
// go는 퍼블릭과 프라이빗이 없습니다.
// 대문자시작은 퍼블릭, 소문자 시작은 프라이빗으로 간주합니다.
func PrintCustom() {
	fmt.Println("사용자 정의 패키지")
}
