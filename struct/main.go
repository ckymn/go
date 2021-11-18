package main

import "fmt"

func main() {

	// var employee struct {
	// 	name      string
	// 	age       int
	// 	isMarried bool
	// }

	// fmt.Println(employee) // { , o , false}

	// employee.age = 22
	// employee.name = "eylem"
	// employee.isMarried = false

	// fmt.Printf("%v \n", employee)

	type employee struct {
		name      string
		age       int
		isMarried bool
		kids      []string
	}

	var e1 employee
	e1.name = "eylem"
	e1.age = 22
	e1.isMarried = false

	fmt.Println(e1)

	// best way dude
	e2 := employee{
		name:      "muhammet",
		age:       22,
		isMarried: false,
		kids:      []string{"eylem", "kubra", "sevgi", "yesim"},
	}

	fmt.Println(e2)
}
