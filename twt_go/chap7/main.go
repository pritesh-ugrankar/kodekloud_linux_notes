package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	x := 5
	val := x <= 5
	fmt.Printf("%t\n", val)
	//Below will print "true" because
	//in ASCII Table, the upper case letters come first and have values like A is 65, B is 66 and so on.
	//The lower case letters have a higher value such as (maybe) 89 or something.
	y := "x"
	z := "A"
	val = y >= z
	fmt.Printf("%t\n", val)
}
