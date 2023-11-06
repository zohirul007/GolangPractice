//If the student receives over 40% on her test, report that her grade passes;
// if not, report that her grade fails.

package main

import "fmt"

func main() {
	grade := 40

	if grade >= 40 {
		fmt.Println("Your grade is: pass")
	} else if grade < 40 {
		fmt.Println("Your grade is: fail")
	}
}
