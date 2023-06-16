package main

import "testing"

func TestSquare1(t *testing.T) {
	result := square(3)
	//테스트 실패
	if result != 9 {
		t.Errorf("기대값이 9인데 실제 값은 %d 입니다.", result)
	}
}
