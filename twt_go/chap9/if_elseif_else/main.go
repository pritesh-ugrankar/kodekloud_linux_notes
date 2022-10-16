package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Printf("Enter your age:\t")
	var ageLimit int64 = 18
	inputAge := bufio.NewScanner(os.Stdin)
	inputAge.Scan()
	convInputAge, err := strconv.ParseInt(inputAge.Text(), 10, 64)

	if err == nil {
		if convInputAge >= ageLimit {
			fmt.Println("You are", convInputAge, "years old so you can ride alone.")

			//You could have written the else if loop as shown below and it  would still have worked.
			//} else if convInputAge <= 17 && convInputAge >= 12 {
			//OR you could write it like below and it would still work.
		} else if convInputAge >= 12 && convInputAge <= 17 {
			fmt.Printf("You are %d years of age and therefore we will need a guardian to ride with you.\n", convInputAge)
		} else {
			fmt.Println("You are less than", convInputAge, "therefore you cannot ride.")
		}

	} else {
		log.Fatal(err)
	}

}
