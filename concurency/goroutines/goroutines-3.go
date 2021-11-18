package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		// burda eger direk i degerini yazdirsak goroutine schedule'i
		// tum degerleri okuyocal ama referans olarak sadece 10 degerini
		// alacagi icin tum degerlere 10 basacak.
		// bu yuzden referans copy islemi yapiyoruz.
		// clousure - scope- referans ....
		go func(val int) {
			fmt.Println(val)
		}(i)
	}

	time.Sleep(time.Second * 3)
}
