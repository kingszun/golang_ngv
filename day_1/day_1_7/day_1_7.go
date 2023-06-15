//go 에서 파일의 내용을 읽어주는 함수
// io/util.ReadFile(파일 경로를 문자열로 대입) (데이터, error)
// 이 함수가 파일을 읽으면 내용을 데이터에 저장하고 error는 nil이 됩니다.
// 파일을 읽는데 실패하면 데이터는 없고 error가 에러가 발생한 이유를 가지고 있음

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 장점 : 뒤에서 b, err 사용 가능
	// 질문 : 과연 뒤에서 err 또 찍을 필요가 있을까?
	b, err := ioutil.ReadFile("./hello.txt")
	if err == nil {
		//읽었을 때 수행할 내용
		fmt.Println(b)
	} else {
		//읽는데 실패했을 때 내용
		fmt.Println(err)
	}

	// go에서는 if 안에서 함수 호출이 가능
	// 함수 호출 결과로 만들어지는 변수는 if 블럭 안에서만 사용 가능
	// 함수가 ㄷㄱ객를 리턴할 때 주로 이 방식을 이용
	if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
		//읽었을 때 수행할 내용
		fmt.Println(b)
	} else {
		fmt.Println(err)
	}
}
