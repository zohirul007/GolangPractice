// If he has money in his account, calculate interest; if he doesn’t,
// charge a penalty fee
// solving with if condition
// To add an else statement

package main

import "fmt"

func main() {
	balance := 500

	if balance < 0 {
		fmt.Println("Balance is below 0, add funds now or you will be charged a penalty.")
	} else {
		fmt.Println("Your balance is 0 or above.")
	}
}
