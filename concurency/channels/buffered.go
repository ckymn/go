package main

import (
	"fmt"
	"time"
)

func main() {
	buffered_channel := make(chan string, 3)

	go func() {
		buffered_channel <- "first"
		fmt.Println("1. sended")
		buffered_channel <- "ikinci"
		fmt.Println("2. sended")
		buffered_channel <- "third"
		fmt.Println("3. sended")
	}()

	<-time.After(time.Second * 3)

	go func() {
		first_read := <-buffered_channel
		fmt.Println("Receiving...")
		fmt.Println(first_read)
		second_read := <-buffered_channel
		fmt.Println(second_read)
		third_read := <-buffered_channel
		fmt.Println(third_read)
	}()
	<-time.After(time.Second * 5)
}
