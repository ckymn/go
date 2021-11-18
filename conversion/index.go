package main

import "fmt"

func main() {
	x := 10
	y := 10.0

	fmt.Printf("%T %v \n", x, y)
	fmt.Printf("%T %v \n", y, y)

	// Type Conversion type(value)
	fmt.Println(x + int(y))
}
