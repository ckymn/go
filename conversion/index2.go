package main

import "fmt"

func main() {
	var x int8 = 127
	var y int16

	y = int16(x)
	fmt.Println(y)
}
