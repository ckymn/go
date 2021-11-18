package main

import "fmt"

// go da kendi olusturdugumuz veri tiplerini defined type uzerine olustururuz
// ! burda employee struct veri tipi degil. Struct uzerine kurulmus Defined bir veri tipi
type employee struct { // struct --> underlying type, employee --> Defined type , Named type
	name      string
	age       int
	isMarried bool
}

type manager struct {
	employee  // call struct
	hasDegree bool
}

func main() {

	e1 := employee{
		name:      "muhammet",
		age:       22,
		isMarried: false,
	}

	fmt.Println(e1)

	m1 := manager{
		employee: employee{
			name:      "eylem",
			age:       28,
			isMarried: false,
		},
		hasDegree: true,
	}

	fmt.Println(m1)

	// another struct identifier [ ANONIM STRUCT ]

	theBoss := struct {
		name  string
		money bool
	}{name: "The boss", money: true}

	fmt.Println(theBoss)
}
