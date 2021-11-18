package main

import "fmt"

// mile float64 uzerine kurulmus defined kendi veri tipimiz.
type mile float64
type kilometer float64

func main() {
	var m1 mile
	m1 = 3.2
	fmt.Println(m1)
	fmt.Printf("%T %v", m1, m1)

	f1 := float64(4.4)

	fmt.Println()

	// m1 => mile veri tipinde
	// f1 => float64 veri tipinde

	fmt.Println(m1 + mile(f1))
	fmt.Println(float64(m1) + f1)

	k1 := kilometer(7.2)
	fmt.Printf("%T, %v", k1, k1)

	// invalid operator . different type
	fmt.Println(m1 + mile(k1))
}
