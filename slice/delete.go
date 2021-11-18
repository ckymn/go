package main

import "fmt"

func main() {
	myslc := []int{0, 1, 2, 3, 4, 5}

	// ilk n elemani silemek [n:]
	// myslc = myslc[2:]
	// son n elemani silmek [:len(myslc) -n]
	myslc = myslc[:len(myslc)-3]
	fmt.Println(myslc)

}
