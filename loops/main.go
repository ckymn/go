package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	i := 0
	for ; i <= 10; i += 3 {
		fmt.Println(i)
	}
	fmt.Println(i)

	a := 1
	if true {
		a++
		fmt.Println(a)
	}
	fmt.Println(a)

}
