package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	/*
	   age := 11

	   	if age >= 18 {
	   		fmt.Println("You can ride the roller coaster because you are", age)
	   	} else {

	   		fmt.Println("You cannot ride the roller coaster as you are", 18-age, "year(s) short.")
	   	}
	*/

	fmt.Println("vim-go")
	fmt.Printf("Enter your age:\t")

	//You have to declare ageLimit at int64 because, for some reason strconv seems to only use int64.
	//So using inputAge, err := strconv.ParseInt(age.Text(), 10, 8) also does not help.
	// So you have to declare ageLimit as int64
	//OR as shown below, you have to use inputAge8bit as a placeholder variable where you convert inputAge to 8 bit integer
	var ageLimit int8 = 18
	age := bufio.NewScanner(os.Stdin)
	age.Scan()
	inputAge, err := strconv.ParseInt(age.Text(), 10, 8)
	inputAge8bit := int8(inputAge)
	if err == nil {
		if inputAge >= 18 {
			println("you are good to go cause you are", inputAge8bit)
		} else {
			fmt.Printf("you are %d years short for permission\n", ageLimit-inputAge8bit)
		}
	} else {
		//fmt.Println("error is", err)
		log.Fatal(err)
	}

}
