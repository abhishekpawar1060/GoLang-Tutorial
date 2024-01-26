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
