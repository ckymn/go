package main

import (
	"fmt"
)

func main() {
	// buffered channel initialization
	buffered := make(chan int, 5)
	// without wait group on channels
	done := make(chan bool)

	go func(bufChan chan int, done chan bool) {
		for val := range bufChan {
			fmt.Println(val)
		}

		fmt.Println("channel closed")
		done <- true
	}(buffered, done)

	buffered <- 1
	buffered <- 2
	buffered <- 3
	buffered <- 4
	buffered <- 5
	buffered <- 6
	buffered <- 7
	buffered <- 8
	buffered <- 9
	close(buffered)

	// wait goroutine to finish
	// *blocking [ eger done'a data gelene kadar burasindan sonrasi bloklanmis durumda]
	<-done

	fmt.Println("main closed")
}
