/*
	We can use go conditional with some syntactic sugar that GO offers

to shorten up code in some cases. For example instead of writing*/

package main

import "fmt"

func main() {

	if balance := -5; balance < 0 {

		fmt.Println("Your balance is below 0, add funds now or you will be charged a penalty.")

	}
}
