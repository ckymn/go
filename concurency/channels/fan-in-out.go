// birden fazla cahnnel 'a input geliyorsa
// bu channel'lari bir araya getirip tek channel olarak disari vermek

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	outputchan := make(chan int, 10000)

	done := make(chan int)

	go func() {
		for {
			fmt.Println("worker waits...")
			// 2.parametre = close or not [true or false]
			val, open := <-outputchan

			if !open {
				break
			}

			fmt.Println(val)
		}
		done <- 1

	}()

	go func() {
		for {
			c1 <- 1
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			c2 <- 2
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			c3 <- 3
			time.Sleep(time.Second)
		}
	}()

	fanIn([]chan int{c1, c2, c3}, outputchan)

	<-done
}

// bu uc channel'i alip tek channel olarak output uretmek
func fanIn(inChans []chan int, outputChan chan int) {
	for _, ch := range inChans {
		go func(c chan int) {
			for {
				val, open := <-c
				if !open {
					break
				}
				outputChan <- val
			}
		}(ch)
	}
}
