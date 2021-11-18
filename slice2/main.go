package main

import "fmt"

func main() {

	var myarr1 [4]int
	fmt.Println(myarr1)

	myarr2 := [4]int{}
	fmt.Println(myarr2)

	var myarr3 = [4]int{}
	fmt.Println(myarr3)

	// ----------------------

	var myslc1 []int
	fmt.Println(myslc1)

	myslc2 := []int{}
	fmt.Println(myslc2)

	var myslc3 = []int{}
	fmt.Println(myslc3)

	var myslc4 = make([]int, 4)
	fmt.Println(myslc4)

}
