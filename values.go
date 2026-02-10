package main

import "fmt"
import "math"
// import (
// 	"fmt"
// 	"math"
// ) -> other type of import(s)

const s string = "constant"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b,c int = 1,2
	fmt.Println(b+c)
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e) // e = 0

	f := "apple"
	fmt.Println(f)

	fmt.Println(s)
	fmt.Println(s + "a")

	const n = 500000000
	const g = 3e20 / n

	fmt.Println(g)
	fmt.Println(int64(g))
	fmt.Println(math.Sin(n)) // sinus (trigono)


	
}