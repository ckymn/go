package main

import "fmt"

func main() {
	arrs := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, arr := range arrs {
		fmt.Println(i, arr*arr)
	}
}
