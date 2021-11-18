package main

import (
	"fmt"
	"sync"
)

func benchmarkGoroutine() {

	wg := sync.WaitGroup{}
	wg.Add(10000)
	lock := sync.Mutex{}
	total := 0
	for i := 0; i < 10000; i++ {
		go func() {
			lock.Lock()
			total += 1
			lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("benchmark goroutine :", total)
}

func benchmarkNormal() {
	total := 0

	for i := 0; i < 10000; i++ {
		total += 1
	}

	fmt.Println("benchmark normal : ", total)
}
