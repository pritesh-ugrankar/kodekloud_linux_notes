package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	aSlice := []int{3, 4, 5}

	fmt.Println("aSlice: ", aSlice)
	bSlice := aSlice

	bSlice[0] = 100

	fmt.Println("aSlice: ", aSlice, "bSlice: ", bSlice)

	aMap := map[string]string{
		"firstName": "Pritesh",
		"midName":   "Manohar",
		"lastName":  "Ugrankar",
	}
	fmt.Println(aMap)

	bMap := aMap
	bMap["firstName"] = "Ganesh"
	for key, value := range aMap {
		fmt.Println(key, value)
	}

}
