package main

import "fmt"

func main() {
	bolum, kalan := division(104, 5)
	fmt.Println(bolum, kalan)
}

func division(bolunen, bolen int) (bolum, kalan int) {
	bolum = bolunen / bolen
	kalan = bolunen % bolen
	return bolum, kalan
}
