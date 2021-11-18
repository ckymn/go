package main

import "fmt"

func main() {
	var name = "muhammet"
	name = "ali"

	fmt.Println(name) // correct

	var surname = "cokyaman"
	//surname := "yaman"

	fmt.Println(surname) // fail [ in here decleartion and assign both not true]

	surname, age := "yaman", 21

	fmt.Println(surname, age) // correct [ in here you should use different another variable]
}
