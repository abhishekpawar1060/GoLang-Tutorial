package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Gorutines")
	go greeter("Hello")
	greeter("World")
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}