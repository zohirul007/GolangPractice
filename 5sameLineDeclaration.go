package main

import "fmt"

func main() {

	age, name := 33,  "is the age of Tanjin Tisha"

	// We can declare multiple variables on the same line in Golang
	// age := 33
	// name := "The name of Tanjin Tisha"
	// is the same as ⬇️
	// age, name := 33,  "is the age of Tanjin Tisha"

	fmt.Println(age, name)
}