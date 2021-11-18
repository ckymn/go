package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := evenNum(10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("girdiginiz sayi : ", result)
	}
}

func evenNum(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("Error: You must enter an even number")
	}

	return num, nil
}
