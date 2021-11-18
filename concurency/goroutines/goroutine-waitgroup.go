package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		// long running task
		fmt.Println("hello from goroutines")
		wg.Done()
	}()

	// blocking
	wg.Wait()

	fmt.Println("hello from main")
}
