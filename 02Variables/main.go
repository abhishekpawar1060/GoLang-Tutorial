package main

import "fmt"

// jwtToken := 4500;  //-> Not allowed outside the function
// var jwtToken = 4500; //-> allowed
// var jwtToken int = 4500 ;//-> allowed

const LoginToken string = "sjd" // -> Public

func main() {
	var userName string = "Abhishek"
	fmt.Println(userName)
	fmt.Printf("Varialbe is of type : %T \n", userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is type of : %T \n", smallVal)

	var smallFloat float64 = 255.5515455656444556
	fmt.Println(smallFloat)
	fmt.Printf("Variable is type of : %T \n", smallFloat)

	//default values & some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	// implicit type
	var website = "google.com"
	fmt.Println(website)

	// no war style
	numberOfUser := 4665
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

}
