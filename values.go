package main

import "fmt"

const s string = "constant"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 even")
	} else {
		fmt.Println("7 odd")
	}

	if num := 9; num < 0 { // initialization in if-else scope
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "is 1 digit")
	} else {
		fmt.Println(num, "is multiple digits")
	}
}