package main

import (
	"fmt"
)

func main() {
	//simple switch
	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// case 4:
	// 	fmt.Println("four")
	// case 5:
	// 	fmt.Println("five")
	// default:
	// 	fmt.Println("other")
	// }

	//multiple condition switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it's weekend")
	// default:
	// 	fmt.Println("it's working day")
	// }

	//type switch [interface{} -> any type of value can be of i]
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("others")
		}
	}

	whoAmI(3.04)
}
