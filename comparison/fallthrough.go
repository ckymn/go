package main

import "fmt"

func main() {

	switch grade := 2; grade {
	case 2:
		fmt.Println("value is 2")
		fallthrough
	case 1:
		fmt.Println("value is 1")
		fallthrough
	case 0:
		fmt.Println("value is 0")
		fallthrough
	default:
		fmt.Println("value is not valid")
		break
	}
}
