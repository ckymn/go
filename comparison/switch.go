package main

import "fmt"

func main() {
	// grade := 2

	// switch grade {
	// case 2:
	// 	fmt.Println("value is 2")
	// case 1:
	// 	fmt.Println("value is 1")
	// case 0:
	// 	fmt.Println("value is 0")
	// default:
	// 	fmt.Println("value is not valid")
	// }

	switch grade := 3; grade {
	case 2:
		fmt.Println("value is 2")
	case 1:
		fmt.Println("value is 1")
	case 0:
		fmt.Println("value is 0")
	default:
		fmt.Println("value is not valid")
	}
}
