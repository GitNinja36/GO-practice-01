package main

func add(a, b int) int {
	return a + b
}

func getLanguage() (string, string, bool) {
	return "golang", "javascript", true
}

func processIt(fn func(a int) int) {
	fn(1)
}

func processIt2() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	// fmt.Println(add(3, 5))

	// fmt.Println(getLanguage())
	//other way is
	// lang1, _ , lang3 := getLanguage()
	// fmt.Println(lang1, lang3)

	// fn := func(a int) int {
	// 	return 2
	// }
	// processIt(fn)

	fn := processIt2()
	fn(6)
}
