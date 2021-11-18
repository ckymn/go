package main

import (
	"fmt"
	"time"
)

func main() {
	heros := []string{"Marvel", "Flash", "Thanos", "Flash", "Hulk", "Thor",
		"Marvel", "Flash", "Thanos", "Flash", "Hulk", "Thor"}

	go findA(heros)
	go findB(heros)
	fmt.Scanln()

}

func findA(arr []string) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == "Flash" {
			fmt.Println("Bulucu A: " + arr[i] + " buldu")
		}
		time.Sleep(time.Second)
	}
}

func findB(arr []string) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == "Flash" {
			fmt.Println("Bulucu B: " + arr[i] + " buldu")
		}
		time.Sleep(time.Second)
	}
}
