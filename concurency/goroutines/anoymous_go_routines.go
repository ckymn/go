package main

import "fmt"

func main() {
	go func() {
		fmt.Println("hello i'm anonymous go routine function")
	}()
}
