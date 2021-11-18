package main

import "fmt"

func main() {
	myslc := []int{1, 2, 3}
	myslc2 := []int{4, 5}

	// burda myslc "int" bir veri tipidir . myslc2 'nin veri tipi "slice int" oldugu icin atama yapamayiz
	//myslc = append(myslc, myslc2)

	// burda calısması ıcın "variadic" func. ayirmalıyız
	myslc = append(myslc, myslc2...)

	fmt.Println(myslc)
}
