package main

import "fmt"

func add(x int, y int) (int, int, int) {
	return x, y, x + y
}
func main() {
	fmt.Println("vim-go")

	//Because you are "return"ing 3 values from the add function,
	//you have to declare three variables, else you'll get an error.
	num1, num2, total := add(5, 7)
	fmt.Println(num1, "+", num2, "=", total)

}
