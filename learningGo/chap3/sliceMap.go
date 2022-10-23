// Main package
package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	/*
		var x = []int{1, 2, 3, 4, 5}
		fmt.Println("x is ", x)

		x = append(x, 6)
		fmt.Println("After the statement x = append(x, 6) is now ", x)
		y := x
		fmt.Printf("Mem addr of x is %p\n", &x)
		fmt.Printf("Mem addr of y is %p\n", &y)
		fmt.Println("After statement y := append(x, 10), the value of x  is still ", x)
		fmt.Println("After statement y := append(x, 10), the value of y is ", y)
		y[0] = 100
		fmt.Println("After statement y[1] = 100, the value of y is ", y)
		fmt.Println("NOTE that after the above statement the value of x changes!! ", x)
		fmt.Printf("Mem addr of x is %p\n", &x)
		fmt.Printf("Mem addr of y is %p\n", &y)
	*/

	/*
		var x = make([]int, 5)
		fmt.Println("x is ", x)
		//x = append(x, 1, 2, 3, 4, 5)
		for index := range x {
			x[index] = 1 + index
			fmt.Println("here it is.", index, index+1)
		}
		fmt.Println("x is now", x)

		var mySlice = []int{1, 2, 3, 4, 5}

		mySliceOfSlice := mySlice[:]
		myNewSlice := append(mySlice)

		fmt.Printf("mySlice address is %p\nmySliceOfSlice address is %p\n", mySlice, mySliceOfSlice)
		fmt.Printf("myNewSlice address is %p\n", myNewSlice)
	*/
	/*
		var x = []int{1, 2, 3, 4}
		var y = make([]int, 4)
		num := copy(y, x)
		fmt.Println(x, y, num)
	*/

	/*
		var s string = "Hello there"
		var s2 string = s[2:4]
		fmt.Printf("Memory address of s is %p and that of s2 string is %p\n", &s, &s2)
	*/

	/*
		var myStringSlice []int
		if myStringSlice == nil {
			fmt.Println("myStringSlice is nil and contains", myStringSlice)
		} else {
			fmt.Println("myStringSlice is NOT nil and contains", myStringSlice)

		}

		fmt.Printf("Variable myStringSlice contains %v and is at %[1]p\n", myStringSlice)
	*/

	/*
		//****************************MAPS****************************
		//This is a map whose keys are string and value is a string.
		var myMap1 = map[string]string{

			"something":  "this is a string",
			"otherthing": "this is another string",
		}

		fmt.Printf("myMap1 is of type %T and contains %[1]v\n", myMap1)

		//This is a map whose keys are string but values are string slices.
		var myMap2 = map[string][]string{
			"fruits": {"apples", "peaches", "oranges"},
			"snacks": {"chiwda", "laadu", "chakli"},
			"drinks": {"milk shake", "sweet lemon juice", "lemonade"},
		}
		fmt.Printf("myMap2 is of type %T and contains %[1]v\n", myMap2)
	*/

	/*
		mapStrNums := map[string][]int{
			"One":   {10, 100, 1000},
			"Two":   {20, 200, 2000},
			"Three": {30, 300, 3000},
		}
		mapStrNums["Three"] = []int{33, 333, 333}
		mapStrNums["Four"] = []int{40, 400, 4000}

		for key, value := range mapStrNums {
			fmt.Println("Before multiplying values by 3 we have", key, value)
			fmt.Printf("Type of the key variable is %T and type of the value variable is %T\n", key, value)
			for eachValue := range value {
				value[eachValue] *= 3
			}
			fmt.Println("AFTER multiplying values by 3 we have", key, value)
			fmt.Println("---------------------------------------------------")
		}

	*/

	/*
		//Example 3.10
		totalWins := map[string]int{}
		totalWins["Orcas"] = 1
		totalWins["Lions"] = 2
		fmt.Println(totalWins["Orcas"])
		fmt.Println(totalWins["Kittens"])
		totalWins["Kittens"]++
		fmt.Println(totalWins["Kittens"])
		totalWins["Lions"] = 3
		fmt.Println(totalWins["Lions"])
	*/

	/*

		myMap := map[string]int{
			"hello": 5,
			"world": 0,
		}

		//Note the difference. If you say
		// for k, v := range myMap {
		//do something here
		//}
		//The k becomes the key and v becomes the value of the map.
		//HOWEVER, if you say the following
		// k, v := myMap["key"]
		//the "k" contains the value and the v contains a "boolean" value (YES, A BOOLEAN VALUE!!)
		//Therefore the above expression is better written as v, ok := myMap["hello"]
		//v, ok := myMap["hello"]
		//fmt.Printf("Value of variable 'v' is %v and type is %[1]T\n", v)
		//fmt.Printf("Value of variable 'ok' is %v and type is %[1]T\n", ok)
			for key := range myMap {
				v, ok := myMap[key]
				fmt.Println("v: ", v, "ok: ", ok)

			}
	*/

	//Example 3.11
	intSet := map[int]bool{}

	vals := []int{22, 5, 10, 2, 5, 8, 22, 7, 3, 9, 1, 2, 0}

	for _, v := range vals {
		intSet[v] = true

	}
	//fmt.Println("intSet is now:", intSet)
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[500] {
		fmt.Println("We have 500 in the intSet.")
	}

}
