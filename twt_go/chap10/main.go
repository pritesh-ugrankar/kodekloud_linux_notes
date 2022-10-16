package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	/*
		count := 0
		for {
			count++
			fmt.Println("Count is now", count)
			if count == 10 {
				break
			}
		}
	*/

	for x := 0; x <= 5; x++ {
		fmt.Println("Printing x:", x)
	}

	for y := 0; y <= 5; y += 2 {
		fmt.Println("Printing y:", y)
	}

	fmt.Printf("\nThe example  for the continue keyword starts here.\n")
	for num := 0; num <= 1000; num++ {
		if num != 0 && num%3 == 0 && num%7 == 0 && num%9 == 0 {

			//A * is printed for a number that is divisible by 3 and 7 and 9.
			fmt.Printf("*\t")
			continue
		} else {
			fmt.Printf("%d\t", num)
		}

	}
	fmt.Printf("\nThe example  for the break keyword starts here.\n")
	for x := 1; x <= 100; x++ {
		fmt.Println(x)
		if x%3 == 0 {
			fmt.Printf("Reached the number %d which is divisible by 3.\n", x)
			fmt.Printf("Will now IMMEDIATELY BREAK out of the for loop.\n")
			break
		}
	}
	fmt.Printf("\nThe example  for the break keyword ends here.\n")

}
