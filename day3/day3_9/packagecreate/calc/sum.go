package calc

import "fmt"

// 이 패키지가 import 될 때 자동호출되는 함수
// main 보다 먼저 호출됩니다.
// 여기서 설정 파일이나 외부 네트워크에서 설정 내용을 읽어서 초기화 하는 경우가 많습니다.
func init() {
	fmt.Println("calc 패키지 사용")
}
func Sum(a int, b int) int {
	return a + b
}
