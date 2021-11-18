package main

import (
	"fmt"
)

func main() {
	city := [...]string{"1", "2", "3", "4"}

	fmt.Println(city)
	fmt.Println(city[0])
	fmt.Println(len(city))

	var myArr [5]int
	fmt.Println(myArr) //00000
	myArr[0] = 100
	myArr[len(myArr)-1] = 200
	fmt.Println(myArr) //00000

}
