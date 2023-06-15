package main

import "fmt"

func printmap(m map[string]int) {
	fmt.Println("printmap")
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	// 문자열을 키로 하고 숫자를 값으로 하는 map
	m := map[string]int{"adam": 1, "soccer": 2}
	printmap(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
	// 데이터 추가 및 갱신
	m["cyver"] = 3
	m["adam"] = 4
	printmap(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
	// 데이터 삭제
	delete(m, "soccer")
	printmap(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
