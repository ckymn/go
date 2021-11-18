package main

import "fmt"

func main() {
	myslc := []int{1, 2, 3}
	myslc2 := make([]int, 2)
	myslc3 := make([]int, 2)

	// if we use assign method
	myslc2 = myslc

	fmt.Println(myslc)
	fmt.Println(myslc2)

	myslc[0] = 100

	fmt.Println(myslc)
	fmt.Println(myslc2)

	// if we use copy method
	copy(myslc3, myslc)

	fmt.Println(myslc)
	fmt.Println(myslc3)

	myslc[1] = 20

	fmt.Println(myslc)  // [100 20 3]
	fmt.Println(myslc3) // [100 2]  don't change

}
