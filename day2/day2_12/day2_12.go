package main

import (
	"fmt"
	//strconv 패키지에 Atoi라는 함수가 존재
	// 문자열을 받아서 정수로 변환해주는 함수
	// func Atoi(s string)(int, error)
	// 리턴되는 int는 정상적으로 수행이 된 결과
	// error는 실패했을 때 메세지
	"errors"
	"strconv"
)

//error 메세지는 왜 만들까
// error는 아니지만 함수 동작의 범위를 제한하기 위해서 만듭니다.

// n의 값이 0~100사이가 아니라면 에러를 리턴
func scoreinput(n int) error {
	if n < 0 || n > 10 {
		return errors.New("점수는 0~100 사이의 정수여야 합니다.")
	} else {
		return nil
	}
}

// panic발생부를 defer와 함께 함수로 가져오면 함수를 종료시키는걸로 끝남
func f() {
	ar := []int{10, 20, 30}
	defer func() {
		// 이 코드 아래서 panic이 발생하면 여기서 처리하고 작업을 계속 수행하도록 합니다.
		recover()
	}()
	fmt.Println(ar[0])
	panic("메세지") // 강제로 패닉 발생
	fmt.Println(ar[1])
	fmt.Println(ar[2])
	// 컴파일 할때는 타입만 확인하기 때문에 에러가 발생하지 않음
	// 실행시 panic 발생
	fmt.Println(ar[3])
}

func main() {
	if result, err := strconv.Atoi("a100a"); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	for i := -5; i < 15; i++ {
		fmt.Println(i)
		if err := scoreinput(i); err == nil {
			fmt.Println("정상수행")
		} else {
			fmt.Println(err)
		}

	}

	f()

	fmt.Println("프로그램 종료")
	ar := []int{10, 20, 30}
	defer func() {
		// 이 코드 아래서 panic이 발생하면 여기서 처리하고 작업을 계속 수행하도록 합니다.
		recover()
	}()
	fmt.Println(ar[0])
	fmt.Println(ar[1])
	fmt.Println(ar[2])
	// 컴파일 할때는 타입만 확인하기 때문에 에러가 발생하지 않음
	// 실행시 panic 발생
	fmt.Println(ar[3])

	fmt.Println("프로그램 종료")
}
