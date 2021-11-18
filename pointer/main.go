// adres operatorleri

package main

import "fmt"

func main() {

	// name := "arin"

	// fmt.Println(name)  // degisken
	// fmt.Println(&name) // degisken adresi

	// fmt.Printf("%T, %v", &name, &name) // *string ( asterek string ), 0xc000010230%

	i, j := 21, 22

	fmt.Println(&i)
	fmt.Println(&j)
	fmt.Println(*(&i))
	fmt.Println(*(&j))

	p := (&i)
	*p = 61
	fmt.Println(i)

}
