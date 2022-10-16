/*
Consider the following array:
someArr := [5]int {1,2,3,4,5}
*/
package main

import "fmt"

func main() {
	someArr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(someArr)
	fmt.Println("Length of someArr is", len(someArr))
	fmt.Println("Capacity of someArr is", cap(someArr))

	//arrSlice := someArr[] does not seem to work.
	var sliceSomeArray []int = someArr[1:3]
	fmt.Println("sliceSomeArray is", sliceSomeArray)
	fmt.Println("Length of sliceSomeArray is", len(sliceSomeArray))
	fmt.Println("Capacity of sliceSomeArray is", cap(sliceSomeArray))
	/*
		This is because the sliceSomeArray starts at element 1 of someArr which is 2
		and has 3 more elements after that, 3, 4 and 5. So in total 4 and therefore
		the capacity of the SLICE is 4.
	*/
	//You can directly declare a slice as shown.
	aNewSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("You can declare a new slice like this: aNewSlice = []int {1,2,3,4,5,6,7,8,}.")
	fmt.Printf("Doing above creates a new slice of type %T and has %[1]v\n", aNewSlice)
	//This is an array
	a := [3]string{"this", "is", "array"}

	//This is a slice.Note how the element numbers are Not mentioned in a slice.
	b := []string{"This", "is", "a", "slice"}
	fmt.Printf("Array \"a\" is of type %T and has %[1]v\n", a)
	fmt.Printf("Array \"b\" is of type %T and has %[1]v\n", b)

	//Use of "make" to create a slice.
	c := make([]float32, 5)
	fmt.Printf("Type of c is %T and value is %[1]v\n", c)

}
