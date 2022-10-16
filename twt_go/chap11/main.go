package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	answer := 14

	//If the "answer" variable is set to 1, -1 or 14
	// the "switch case" will print BOTH the lines.
	switch answer {
	case 1, -1, 14:
		fmt.Println("1. +one")
		fmt.Println("2. -one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("Could not find the case.")
	}
}
