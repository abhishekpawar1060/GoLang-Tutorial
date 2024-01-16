package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// no inheritance in golang; No super or parent

	abhishek := User{"Abhishek", "abhi@gmail.com", true, 16}
	fmt.Println(abhishek)

	fmt.Printf("Abhishek details are: %+v", abhishek)

	fmt.Printf("Name is %v and email is %v.", abhishek.Name, abhishek.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
