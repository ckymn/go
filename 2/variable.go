package main

import "fmt"

func main() {
	var name string = "muhammmet" //static variable type
	var age int = 22
	surname := "cokyaman" //dinamik variable type // short declaration
	var isMarried bool = false
	var height int      // zero values
	var language string // zero values
	var isgo bool

	fmt.Println(name, surname, age, isMarried)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", surname)
	fmt.Printf("%T\n", age)

	fmt.Println(height)   // 0
	fmt.Println(language) // ""
	fmt.Println(isgo)     // false
}
