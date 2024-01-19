package main

import "fmt"

func main() {
	fmt.Println("Method in golang")

	abhishek := User{"Abhisek", "abhi@gmailcom", true, 21}

	abhishek.GetStatus()

	abhishek.NewMail() //this not change the original email
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "ab@gmail.com"
	fmt.Println("Email of the user is:", u.Email)
}
