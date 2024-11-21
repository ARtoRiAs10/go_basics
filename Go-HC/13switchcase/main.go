package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 6
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	// Use a switch case to handle different dice values
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1, and you can open")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll the dice again")
	default:
		fmt.Println("What was that!")
	}
}
