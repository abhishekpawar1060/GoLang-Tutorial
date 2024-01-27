package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Create Json data in GoLang")

	EncodeJsom()

	DecodeJson()
}

func EncodeJsom() {
	lcoCourses := []course{
		{"React", 2500, "learnCodeOnline.com", "abc123", []string{"web-dev", "JS"}},
		{"GO", 5000, "freeCodecamp.com", "ab123", []string{"backend-Dev", "Go"}},
		{"MERN", 1500, "learnCodeOnline.com0", "54c123", nil},
	}

	//package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "React",
		"Price": 2500,
		"website": "learnCodeOnline.com",
		"_": "abc123",
		"tags": ["web-dev","JS"]
	}	
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	//Some cases where you just want to add data to key value
	fmt.Println("")
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	fmt.Println("")

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and Type is : %T \n", k, v, v)
	}

}
