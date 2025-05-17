package main

import "fmt"

// range -> iterating over data structure
func main() {
	// nums := []int{1, 2, 3}

	//iteration with for loop
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// //iteration with range
	// sum := 0
	// for _, num := range nums {
	// 	// fmt.Println(num)
	// 	sum = sum + num
	// }
	// fmt.Println(sum)

	//iteration with range
	// for i, num := range nums {
	// 	fmt.Println(i, " -> ", num)
	// }

	// m := map[string]string{"fname": "rohit", "mname": "kumar", "lname": "sharma"}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	//unicode pont rune
	// k -> starting byte of rune
	for k, v := range "golang" {
		fmt.Println(k, " -> ", string(v))
		fmt.Println(k, " -> ", v)
		fmt.Println()
	}
}
