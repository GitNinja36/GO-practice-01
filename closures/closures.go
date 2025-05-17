package main

import "fmt"

func closures() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := closures()
	fmt.Println(increment())
}
