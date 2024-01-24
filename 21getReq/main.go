package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("get web Request in go lang")

	getReq()
}

func getReq() {
	const myURL = "http://localhost:7000/get"

	response, err := http.Get(myURL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code : ", response.StatusCode)
	fmt.Println("Content length is : ", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is : ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}
