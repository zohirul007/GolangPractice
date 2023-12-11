/*
	We can use go conditional with some syntactic sugar that GO offers

to shorten up code in some cases. For example instead of writing

package main

	 import "fmt"

	 func main () {
		grade := 89

		if grade >= 90 {
			fmt.Println("A grade")
*/
package main

import "fmt"

func main() {
	if grade := 89; grade < 90 {
		fmt.Println("B grade")
	}

}
