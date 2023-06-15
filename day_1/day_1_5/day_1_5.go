package main

import (
	"fmt"
	"os"
)

// go run <binary_name> arg1 arg2 ...
func main() {
	//saving command line args in variables
	arguements := os.Args
	//Printing index and value loop
	for n, arg := range arguements {
		fmt.Println(n, arg)
	}
	//arguments number 0 is binaryfile path
	fmt.Println(arguements[0])

	var x float64 = 10.0
	for i := 0; i < 10; i++ {
		x = x - 0.1
	}
	fmt.Println(x)
	fmt.Println(x == 9.0)

	ch := 'a' // 1 literature - byte
	fmt.Printf("%c \n", ch)
	fmt.Println(ch)

	unicode := 'd' // 1 unicode - rune
	fmt.Printf("%c\n", unicode)
	fmt.Println(unicode)

	name := "park"
	fmt.Println(name)

	var gender string = "민준"
	fmt.Println(gender)
}
