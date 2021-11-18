// blocking type

// her okuma yazma islemi blocking'tir
// sadece es zamanli olarak bir data yazip bir data okuyabilirsin
package main

import (
	"fmt"
)

func main() {
	unbuffered := make(chan int)
	var unbuffered2 chan int // must use make() func.

	fmt.Println(unbuffered)
	fmt.Println(unbuffered2) // nil

	// 1. oncelikle go schedule edilir.
	go func(unbufChan <-chan int) {
		// 2. ikincisi burda blocking olur. veri gelene kadar
		value := <-unbufChan
		fmt.Println(value)
	}(unbuffered)

	// veri chanel'a gidince blocking open and goroutine calisir.
	unbuffered <- 1
}
