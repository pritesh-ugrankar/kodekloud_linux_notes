package main

import "fmt"

func main() {
	/*
		Note, this program will "run" and "build" just fine.
		But it will "panic" and exit during runtime because
		you cannot "divide by zero".
	*/
	var numone int = 8
	var numtwo int = 0
	oneandtwo := numone / numtwo
	println("Answer is ", oneandtwo)
	var num1 float64 = 7
	var num2 int = 5
	answer := num1 / float64(num2)
	fmt.Println("Answer is ", answer)
	fmt.Printf("Answer is %08.4f\n", answer)
}
