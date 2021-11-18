package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("hello from goroutines")
	}()

	// eger bunu yazmaz isek goroutines calismaz cunku
	// burda calismaya hazir hale getirip calistirip cikicak vakti olamadi
	// burda kesinlikle yazacak diyemeyiz cunku cok hizli calisiyor.
	// ! Bundan kurtulmak icin waitGroup primitive tipleri kullanabiliriz.

	fmt.Println("hello from main")
}
