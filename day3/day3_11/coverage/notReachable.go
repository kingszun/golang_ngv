package notReachable

import "fmt"

func S1() {
	fmt.Println("Hi")
	return
	fmt.Println("?????")
}

func S2() {
	return
	fmt.Println(">>>>>>/?????")
}
