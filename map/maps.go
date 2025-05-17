package main

import (
	"fmt"
	"maps"
)

// maps -> object, hash, dictionary
func main() {
	// // creating maps
	// m := make(map[string]string)

	// //setting an element nameOfMap["key"] = "value"
	// m["name"] = "golang"

	// // get element
	// // IMP : if key doesn't exist in the map it will return zero value
	// fmt.Println(m["name"])

	//to get the length of the map
	// fmt.Println(len(m))

	//delete
	// delete(m, "name")

	//to cler the map
	// clear(m)

	// m := make(map[string]int)
	// m["age"] = 30
	// fmt.Println(m["age"])
	// fmt.Println(len(m))
	// delete(m, "age")
	// fmt.Println(len(m))

	// m := map[string]int{"price": 10, "age": 22}
	// fmt.Println(m)

	// k, ok := m["price"]
	// fmt.Println(k)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not all ok")
	// }

	m1 := map[string]int{"price": 10, "age": 22}
	m2 := map[string]int{"price": 10, "age": 22}
	fmt.Println(maps.Equal(m1, m2))

}
