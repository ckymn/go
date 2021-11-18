package main

import "fmt"

type mile float64
type kilometer float64

func main() {

	m1 := mile(160)
	res := mileToKilometer(m1)

	fmt.Println(res)
}

func mileToKilometer(m mile) kilometer {
	return kilometer(m * 1.6)
}
