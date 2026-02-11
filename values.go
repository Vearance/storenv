package main

import "fmt"
import "time"

func main() {
	i := 2

	fmt.Print("Write ", i, " as ")
	switch i {
	case 1: // if i == 1
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:// multiple expression
		fmt.Println("weekend")
	default: // acsts like otherwise
		fmt.Println("weekday")
	}


	t := time.Now()
	switch { //switch without an expression -> if-else logic
	case t.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}


	// compare types - never seen this or im just dumb - in Python its isInstance
	whatAmI := func (i interface{})  { // syntax func
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")


}