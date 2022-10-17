package main

import "fmt"

// test2 is a function that takes a fuction whose parameter is an int and returns int
func test2(myFunc func(int) int) {
	fmt.Println(myFunc(7))

}
func main() {
	fmt.Println("vim-go")

	//Anonymous function
	//Note - test is the var name NOT func name!!
	test := func(x int) int {
		return x * -1
	}

	test2(test)
	test2(test)

}
