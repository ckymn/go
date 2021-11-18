package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		if i%3 == 0 {
			continue // asagi inme tekrar basa don demek
		}

		fmt.Println(i)
	}
}
