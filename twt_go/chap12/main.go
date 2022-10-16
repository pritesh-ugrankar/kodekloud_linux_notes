package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	//This is one way to declare an array.
	var arr [5]int
	fmt.Println("All the elements of the array named arr are set to 0.", arr)
	for i := 0; i < len(arr); i++ {
		//You cannot do var arr [..]int or var arr [...]int
		//because arrays ALWAYS have a predeclared size.

		//You cannot simply do arr[i] = i * 100
		//because then, the first (0th) element
		//will stay 0. You cannot do arr[i] = i + 100
		//because then you will get 101,102..105
		//So you have to do the following.

		arr[i] += (i + 1) * 100
	}

	fmt.Println("After the for loop runs, this is how the \"arr\" array becomes:", arr)
	//This is another way to declare an array.
	newarr := [3]int{4, 5, 6}
	fmt.Println("newarr := [3]int{4, 5, 6} = ", newarr)

	fmt.Println("Multidimensional array")
	arr2D := [5][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}}
	//fmt.Printf("arr2D is of type %T, and contains %[1]v\n", arr2D)
	//This is the only way a two variable initialization at least partially works in go!!
	//for i, j := 0, 0; i < len(arr2D[i]) && j < len(arr2D[j]); i, j = i+1, j+1 {
	//}

	//Note - you have to put 5 and 3. You cannot use len(arr2D[i]) and len(arr2D[j])

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("arr2D[%d][%d] = %d\n", i, j, arr2D[i][j])
		}
	}

}
