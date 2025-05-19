package main

import "fmt"

// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// to allow any value
// func printSlice[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//to allow only perticular value like int and string
// func printSlice[T int | string](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// LIFO
type stack[T any] struct {
	elements []T
}

func main() {
	myStack := stack[string]{
		elements: []string{"golang"},
	}
	fmt.Println(myStack)

	// nums := []int{1, 2, 3}
	// printSlice(nums)
	// names := []string{"golang", "typescript"}
	// printStringSlice(names)

	// nums := []int{1, 2, 3}
	// names := []string{"golang", "typescript"}
	// printSlice(nums)
	// printSlice(names)

}
