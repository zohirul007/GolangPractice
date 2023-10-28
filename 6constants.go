package main

//Constants can be character, string, boolean, or numeric values.
//They can not be more complex types like slices, maps and structs
import "fmt"

func main() {

	const Name = "Tanjin Tisha"
	const Role = "is a Model and Tv actress"

	fmt.Println("what is the name: ", Name)
	fmt.Println("what is she doing: ", Role)
	//As the name implies, the value of a constant can't be changed after it has been declared
}
