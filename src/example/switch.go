package main

import (
	"fmt"
	"time"
)

/**
in go, switch can use to catch type
*/
func main() {
	i := 2
	fmt.Print("write ", i, " as ")
	// catch value
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	// catch multiple value at one case statement
	case time.Sunday, time.Saturday:
		fmt.Println("today is weekend day")
	default:
		fmt.Println("today is work day")
	}

	t := time.Now()
	// switch can work without expression
	// now switch is same with if/else structure somehow
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	// function is a type
	whatAmI := func(i interface{}) {
		// switch can catch type value
		switch t := i.(type) {
		case int:
			fmt.Println(i, "is a int")
		case bool:
			fmt.Println(i, "is a bool")
		case string:
			fmt.Println(i, "is a string")
		default:
			fmt.Printf("%v is a %T type handled in default\n", i, t)
		}
	}
	// use function
	whatAmI(2)
	whatAmI(2.2)
	whatAmI(false)
	whatAmI("china")
}
