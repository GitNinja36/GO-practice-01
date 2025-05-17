package main

import (
	"fmt"
)

func main() {
	//uninitialised slice is null
	// var nums []int
	// fmt.Println(nums)
	// fmt.Println(len(nums))

	//make() -> its a inbuild function were you have to pass 2 to 3 variable the 1st will be the array type, 2nd will be initial size and 3rd will be intial capacity
	// var nums = make([]int, 0, 5)
	// capacity -> maximum number of element can be fit
	// fmt.Println(cap(nums))
	// // fmt.Println(nums)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// fmt.Println(cap(nums))
	// fmt.Println(nums)

	// other way is
	// nums := []int{}
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))
	// fmt.Println(nums)

	// append at perticular place
	// var nums = make([]int, 2, 5)
	// nums[0] = 3
	// fmt.Println(cap(nums))
	// fmt.Println(nums)

	//for copying the function
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// // copy function
	// copy(nums2, nums)

	// fmt.Println(nums)
	// fmt.Println(nums, nums2)

	// //sclice operator
	// var nums = []int{1, 2, 3}
	// // nums[from : to]
	// fmt.Println(nums[0:1])

	// slices
	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 3}
	// fmt.Println(slices.Equal(nums1, nums2))

	//2d array
	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)
}
