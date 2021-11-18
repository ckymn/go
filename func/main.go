package main

import "fmt"

func main() {
	fmt.Println(result(45)) // arguman
}
func result(grade int) string { // parametre
	if grade >= 50 {
		return "gectiniz"
	}
	// else yazmamiza gerek yok artik . return ile bitirebiliriz.
	return "gecemediniz"
}
