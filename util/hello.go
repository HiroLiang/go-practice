package util

import "fmt"

func Hello() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s! \n", s)

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}
}
