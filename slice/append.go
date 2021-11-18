package main

import "fmt"

func main() {
	myslc := []int{1, 2, 3}

	//myslc = append(myslc, 4, 5) // eger burayi acar isek ayni UnderlineArray 'a bakacaklari icin ikisinin 0'indexteki degerleri degisir.

	fmt.Println(myslc)

	myslc2 := append(myslc, 6)

	fmt.Println(myslc2)

	myslc[0] = 100

	fmt.Println(myslc)
	fmt.Println(myslc2)
}
