package main

import "fmt"

type rectangle struct {
	a, b float64
}

func (r rectangle) area() float64 {
	return r.a * r.b
}

func (r rectangle) circumference() float64 {
	return 2 * (r.a * r.b)
}

type shape interface {
	area() float64
	circumference() float64
}

func interfaceFunc(i shape) {
	fmt.Println(i)
	fmt.Println(i.area())
	fmt.Println(i.circumference())
	fmt.Printf("%T", i)
	fmt.Println()
}
func main() {

	r := rectangle{3,8}
	fmt.Println("Area :", r.area())
	fmt.Println("Circumference :", r.circumference())

	interfaceFunc()
}
