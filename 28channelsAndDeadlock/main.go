package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels And Deadlock")

	myCh := make(chan int, 2)

	wg := &sync.WaitGroup{}

	wg.Add(2)

	//Receive Only
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChanelOpen := <-myCh

		fmt.Println(isChanelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	//Send Only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		// myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
