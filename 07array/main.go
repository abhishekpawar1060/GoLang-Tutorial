package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var arr [4]string

	arr[0] = "Zero"
	arr[1] = "One"
	arr[3] = "three"

	fmt.Println("Array list is : ", arr)
	fmt.Println("Array list is : ", len(arr))

	var myArray = [3]string{"Orange", "Apple", "Banana"}
	fmt.Println("Array list is : ", myArray)
	fmt.Println("Array list is : ", len(myArray))

}
