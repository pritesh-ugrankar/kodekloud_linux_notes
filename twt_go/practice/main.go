package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	var myMapStringNums = map[string][]int{}
	myMapStringNums["one"] = []int{1, 10, 100, 1000}
	myMapStringNums["two"] = []int{2, 20, 200, 2000}

	for key, value := range myMapStringNums {
		fmt.Println("-------------------------------")
		fmt.Println("Before: ", key, value)
		for eachValue := range value {
			value[eachValue] *= 100
		}
		fmt.Println("After: ", key, value)
	}
}
