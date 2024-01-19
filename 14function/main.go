package main

import (
	"fmt"
)

func main() {
	fmt.Println("Function in golang")
	myfunc()

	result := sum(45, 50)
	fmt.Println("Sum is", result)

	preRes, myMSG := proAdder(2, 4, 5, 6)
	fmt.Println("Sum is", preRes)
	fmt.Println("Pro Message is:", myMSG)
}

func myfunc() {
	fmt.Println("My function")
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi pro result function"
}
