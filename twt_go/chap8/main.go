package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	val := true && false
	fmt.Println("true & false is", val)
	val = false && true
	fmt.Println("false & true  is", val)
	val = false && false
	fmt.Println("false & false is", val)
	val = true || false
	fmt.Println("true || false is", val)
	val = false || true
	fmt.Println("false || true  is", val)

	//If it is a complex expression like what is shown below
	// remember that if there is a "false" immediately before
	//or after an "&&", then the expression will be false.

	val = (8 > 9) && ("a" == "a") || (10 > 9) || (7 > 23)
	fmt.Println("The expression (8 > 9) && (a == a) || (10 > 9) || (7 > 23) will evaluate to", val)

	val = ((8 > 9) && ("a" == "a") && (10 > 9)) || (42 > 23)
	fmt.Println("val = ((8 > 9) && (a == a) && (10 > 9)) || (42 > 23) will print", val)

	val = ((8 > 9) && ("a" == "a") && (10 > 9)) && (42 > 23)
	fmt.Println("val = ((8 > 9) && (a == a) && (10 > 9)) && (42 > 23) will print", val)

}
