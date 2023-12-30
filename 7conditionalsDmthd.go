// Find out height with conditional on go
/*We can write conditional with some short cut way like this*/
package main

import "fmt"

func main() {
	
	if height := 7; height > 6 {
		fmt.Println("You are super tall!")
	} else if height > 4 {
		fmt.Println("You are tall enough!")
	} else {
		fmt.Println("You are not tall enough!")
		
	}
}
