package main

import "fmt"

func main() {
	//by defalut 0 value will be added
	// string -> "", int -> 0, bool -> false
	// var nums [4]int

	// array length
	// fmt.Println(len(nums))

	//input value
	// nums[0] = 2
	// fmt.Println(nums[0])

	//print full array
	// fmt.Println(nums)

	//boolean values
	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)

	//string values
	var names [4]string
	names[1] = "golang"
	fmt.Println(names)

	// nums := [3]int{1, 2, 3}
	// fmt.Println(nums)

	//2d array
	nums := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(nums)

}
