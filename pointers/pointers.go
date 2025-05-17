package main

import "fmt"

// by Value
func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum : ", num)
}

// by reference
func changeNum2(num *int) {
	*num = 5 //dereference
	fmt.Println("In changeNum : ", *num)
}

func main() {

	num := 1
	// changeNum(num)
	changeNum2(&num)

	// fmt.Println("memory address : ", &num)
	fmt.Println("after changingNum in main : ", num)
}
