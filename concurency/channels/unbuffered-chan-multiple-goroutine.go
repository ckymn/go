package main

import (
	"fmt"
	"time"
)

func main() {
	// channel initialization
	unbufferedChan := make(chan int)

	// reader goroutine
	go func(unBufChan chan int) {
		// * blocking
		value := <-unBufChan
		fmt.Println(value)
	}(unbufferedChan)

	// writer goroutine
	go func(unbufChan chan int) {
		// * blocking open
		unbufChan <- 1
	}(unbufferedChan)

	fmt.Println("hello channels")
	time.Sleep(time.Second * 3)
}
