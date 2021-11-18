package main

import (
	"fmt"
	"time"
)

func main() {
	heros := []string{"Marvel", "Flash", "Thanos", "Flash"}
	firstChannel := make(chan string)

	go func(arr []string) {
		for _, hero := range arr {
			firstChannel <- hero
			time.Sleep(time.Second)
		}
	}(heros)

	go func() {
		for i := 0; i < 4; i++ {
			finded := <-firstChannel
			fmt.Println("Alıcı: Bulucudan " + finded + " alındı")
		}
	}()
	<-time.After(time.Second * 5)
}
