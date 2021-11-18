package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	arr_2 := []int{}
	arr_2 = make([]int, 3)

	fmt.Println(arr)
	fmt.Println(arr_2)
	arr[0] = 5
	arr_2[0] = 12
	fmt.Println(arr)
	fmt.Println(arr_2)

}
