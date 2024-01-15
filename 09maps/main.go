package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languges := make(map[string]string)

	languges["JS"] = "JavaScript"
	languges["RB"] = "Ruby"
	languges["py"] = "Python"

	fmt.Println("List of all languages: ", languges)

	fmt.Println("JS short for : ", languges["JS"])

	delete(languges, "RB")
	fmt.Println("List of all languages: ", languges)

	//loops on maps
	for key, value := range languges {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}

	//other method
	for _, value := range languges {
		fmt.Printf("For key v, value is %v \n", value)
	}

}
