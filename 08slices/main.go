package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Tomato"}
	fmt.Printf("Type of fruitList is %T ", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[:1])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 250
	highScores[1] = 942
	highScores[2] = 132
	highScores[3] = 756
	//highScores[4] = 888

	highScores = append(highScores, 666, 545, 333)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)

	//remove value from slices based on index
	var courses = []string{"JS", "JAVA", "ruby", "PHP"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
