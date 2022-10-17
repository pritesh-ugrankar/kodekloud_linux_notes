package main

import "fmt"

func add(x, y int) (plus, minus, mult int) {
	//return x + y, x - y, x * y
	mult = x * y
	minus = x - y
	plus = x + y
	return
}
func main() {
	fmt.Println("vim-go")

	//Because you are "return"ing 3 values from the add function,
	//you have to declare three variables, else you'll get an error.
	minus, mult, plus := add(4, 3)
	fmt.Println(minus, mult, plus)

}
