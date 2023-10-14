package main

//Boot.dev problem solving
//Textio users are reporting that we're billing them for wildly inaccurate amounts. They're supposted to be billed .02 dollars for each text message sent.

//Fix the the math bug on line 17

import "fmt"

func main() {
	messagesFromDoris := []string{
		"You doing anything later??",
		"Did you get my last message?",
		"Don't leave me hanging...",
		"Please respond I'm lonely!",
	}
	numMessages := float64(len(messagesFromDoris))
	costPerMessage := .02

	// don't touch above this line

	totalCost := costPerMessage * numMessages

	// Print statement

	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
}