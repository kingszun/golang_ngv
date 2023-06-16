package quick

import "testing"
import "testing/quick" // 랜덤하게 기본값을 생성해서 테스트하도록 도와주는 패키지

// 테스트 횟수
var N = 1000000

// 테스트 함수
func TestQuick(t *testing.T) {
	// 테스트 조건:ADD의 매개변수를 좌우 변경해서 수행
	condition := func(a, b Point2D) bool {
		return Add(a, b) == Add(b, b)
	}
	//랜덤한 값을 대입해서 테스트
	if err := quick.Check(condition, &quick.Config{MaxCount: N}); err != nil {
		//에러가 발생한 경우 출력
		t.Errorf("Error:%v\n", err)
	}
}
