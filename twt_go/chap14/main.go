package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	aSlice := []int{4, 1, 3, 4, 56, 7, 12, 4, 6}

	for index, element := range aSlice {
		fmt.Println(index, element)
	}

	/*

		for index, element := range aSlice {
			for j := index + 1; j < len(aSlice); j++ {
				element2 := aSlice[j]

				if element2 == element {
					fmt.Println(element, index)
					break
				}
			}
		}
	*/

}
