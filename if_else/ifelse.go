package main

import "fmt"

func main() {
	// age := 12

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is an teenager")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	var role = "admin"
	var hasPermission = true
	if role == "admin" || hasPermission {
		fmt.Println("yes")
	}

	//go does not ahve ternary operator
}
