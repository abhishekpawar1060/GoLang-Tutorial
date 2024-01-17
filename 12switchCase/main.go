package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is one and you can open")
	case 2:
		fmt.Println("Dice can to move 2 spot")
	case 3:
		fmt.Println("Dice can to move 3 spot")
	case 4:
		fmt.Println("Dice can to move 4 spot")
		//fallthrough
	case 5:
		fmt.Println("Dice can to move 5 spot")
	case 6:
		fmt.Println("Dice can to move 6 spot")

	default:
		fmt.Println("Default")
	}
}
