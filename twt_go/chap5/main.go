package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Enter your birth year: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	/*
		You may want to write the above line as
		input := int (scanner.Text)
		But you cannot do that because you can convert int to float or vice versa
		but not string to int or float.
	*/
	if err == nil {
		fmt.Printf("You are %d years old in 2022\n", 2022-input)
	} else {
		fmt.Printf("%s\n", err)
	}

}
