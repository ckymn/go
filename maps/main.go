package main

import "fmt"

func main() {
	// create map
	myMap := map[string]int{
		"ahmet": 21,
		"ayse":  22,
		"selim": 12,
	}

	// loop with foreach on map
	for i, m := range myMap {
		fmt.Println(i, m)
	}

	// create map with make
	myMap2 := make(map[int]string)
	myMap2[0] = "seher"
	myMap2[1] = "ali"
	fmt.Println(myMap2)

	// condition
	_, ok := myMap2[3]
	fmt.Println(ok) // false

	//delete map
	delete(myMap, "ahmet")
	fmt.Println(myMap)
}
