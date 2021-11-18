package main

import (
	"fmt"
	"sync"
)

type RaceTest struct {
	Val int
}

func main() {
	// type'tan pointer tipinde bir degisken olusturuyor
	raceTest := &RaceTest{}

	wg := &sync.WaitGroup{}
	wg.Add(10000)

	lock := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		lock.Lock()
		go increment(raceTest, wg)
		lock.Unlock()
	}

	wg.Wait()
	fmt.Println("race test :", raceTest)
}

func increment(rt *RaceTest, wg *sync.WaitGroup) {
	rt.Val += 1

	wg.Done()
}
