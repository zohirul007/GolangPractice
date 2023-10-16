package main

import "fmt"

func main() {

	fmt.Print("Please enter first number: ")
	var n1 int
	fmt.Scanln(& n1) //Take input from user
	fmt.Print("Please enter Second number: ")
	var n2 int
	fmt.Scanln(& n2) //Take input from user
	result := n1 + n2
	fmt.Println("Sum of two numbers: ", result)

}
