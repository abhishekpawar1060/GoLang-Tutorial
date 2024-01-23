package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=fefl232ds"

func main() {
	fmt.Println("Welcome to handling URLs in Golang")
	fmt.Println(myURL)

	//parsing
	result, _ := url.Parse(myURL)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are : %T \n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("param is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=abhi",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
