package main

import "fmt"

const s string = "constant"

func main() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	fmt.Println(0)
	fmt.Print(1)
	fmt.Printf("2")

	for {
		fmt.Println("loop")
		break
	}

	for n:= range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}