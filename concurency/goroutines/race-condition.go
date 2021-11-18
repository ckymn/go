package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//raceExample()
	//raceExampleFixed()
	raceExampleFixedWithAtomic()
}
func raceExample() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	//shared value
	val := 0

	// Burda bize 2000 sonucunu vermesi lazim ama vermiyor
	// Nedeni ise race-condition , burda fonc'lar ayni alana ulasip degerleri arttirmaya calisiyorlar
	// 1.func 10 yaparken 2.func 9'dan geliyor mesela.
	go func() {
		for i := 0; i < 1000; i++ {
			val++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			val++
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(val)
}

func raceExampleFixed() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	//mutex : Bir memory adresine es zamanli olarak sadece bir goroutine'nin erisebilecegini garantiler
	lock := sync.Mutex{}

	//shared value
	val := 0

	// first class citizens
	x := func() {
		for i := 0; i < 1000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}
		wg.Done()
	}

	go x()
	go x()

	wg.Wait()

	fmt.Println(val)
}

func raceExampleFixedWithAtomic() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	var val int32 = 0

	// first class citizens
	x := func() {
		for i := 0; i < 1000; i++ {
			// referans adresini ve ne kadar arttirmamiz gerektigini yazariz
			atomic.AddInt32(&val, 1)
		}
		wg.Done()
	}

	go x()
	go x()

	wg.Wait()

	fmt.Println(val)
}
