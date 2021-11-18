package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		if i == 3 {
			break // donguyu kir ve biitr.
		}

		fmt.Println(i)
	}
}
