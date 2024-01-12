package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Sum of Two number by using usefr intput")

	num1 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the value of num1 : ")
	input1, _ := num1.ReadString('\n')

	num2 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the value of num2 : ")
	input2, _ := num2.ReadString('\n')

	fmt.Printf("The sum of num1 and num2 is %v ", input1+input2)

}
