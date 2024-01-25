package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Post request in go lang")

	postJSONReq()

	posyFormReq()

}

func postJSONReq() {
	const muUrl = "http://localhost:7000/post"

	//fake json payload

	reqBody := strings.NewReader(`
		{
			"coursename" : "Lets go with golang",
			"price" : 0,
			"platform" : "learnCodeOnline.in"
		}
	`)

	response, err := http.Post(muUrl, "application/json", reqBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func posyFormReq() {
	const muUrl = "http://localhost:7000/post"

	//formdata
	data := url.Values{}

	data.Add("firstname", "abhi")
	data.Add("lastname", "pawar")
	data.Add("email", "abhi@gmail.com")

	response, err := http.PostForm(muUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
