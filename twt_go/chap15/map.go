package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	var myMap map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	/*
		Another way of declaring a map.
			myMap := map[string]int{
				"apple":  5,
				"pear":   6,
				"orange": 9,
			}
	*/

	fmt.Printf("The variable \"myMap\" is of type %T and contains %[1]v\n", myMap)
	//Access a particular value at a key.
	fmt.Println(myMap["apple"])
	//Change the value of a key.
	myMap["apple"] = 900
	fmt.Println(myMap["apple"])

	//Add a new key and value to an existing map.
	myMap["grapes"] = 200

	//Access key and value both.
	for key, value := range myMap {
		fmt.Println("Key is", key, "and value is", value)
	}

	//Delete a key and value pair.
	delete(myMap, "grapes")
	fmt.Println("Now \"grapes\" is deleted and we get the rest of the fruits:", myMap)
}
