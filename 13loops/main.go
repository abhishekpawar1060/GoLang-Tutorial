package main

import "fmt"

func main() {
	fmt.Println("Loops in goLang")
	count := 10
	for i := 1; i < count; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday", "Saturdy"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and values is %v\n", index, day)
	}

	for _, day := range days {
		fmt.Printf("day is %v\n", day)
	}

	rougheValue := 1
	for rougheValue < 10 {

		if rougheValue == 2 {
			goto lco
		}

		if rougheValue == 5 {
			rougheValue++
			continue
		}
		fmt.Println("Value is:", rougheValue)
		rougheValue++
	}

lco:
	fmt.Println("Jumping at LearnCodeonline.in")

}
